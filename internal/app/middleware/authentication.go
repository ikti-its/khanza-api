package middleware

import (
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func get_private_ip() string {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}
			ip := ipNet.IP
			return ip.String()
		}
	}
	return ""
}

func get_mac_address(ctx *fiber.Ctx) string {
	ip := ctx.IP()
	if ip == "" {
		log.Println("⚠️ Empty IP address, skipping MAC detection")
		return ""
	}

	out, err := exec.Command("arping", "-c", "1", ip).Output()
	if err != nil {
		log.Printf("⚠️ arping command failed for IP %s: %v", ip, err)
		return ""
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Unicast reply") {
			parts := strings.Fields(line)
			if len(parts) >= 5 {
				return parts[4]
			}
		}
	}
	return ""
}

func get_self_mac_address() string {
	iface, err := net.InterfaceByName("wlp0s20f3") // adjust if needed
	if err != nil {
		log.Printf("⚠️ get_self_mac_address failed: %v", err)
		return ""
	}
	return iface.HardwareAddr.String()
}

func get_mac_ip_address(ctx *fiber.Ctx) (string, string) {
	ip := ctx.IP()
	mac := get_mac_address(ctx)

	// Handle localhost
	if ip == "127.0.0.1" || strings.HasPrefix(ip, "::1") {
		ip = get_private_ip()
		mac = get_self_mac_address()
	}

	// ✅ Fallback if MAC is still empty
	if mac == "" {
		mac = "00:00:00:00:00:00" // or optionally: get_self_mac_address()
	}

	return ip, mac
}

func Authenticate(roles []int) func(ctx *fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")

	return jwtware.New(jwtware.Config{
		SuccessHandler: func(ctx *fiber.Ctx) error {
			token, ok := ctx.Locals("jwt").(*jwt.Token)
			if !ok {
				log.Println("❌ JWT token not found in context")
				return fiber.ErrUnauthorized
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Println("❌ JWT claims are invalid")
				return fiber.ErrUnauthorized
			}

			// Extract and assert JWT claims
			sub, okSub := claims["sub"].(string) // ✅ ensure it's a string (UUID)
			roleFloat, okRole := claims["role"].(float64)

			if !okSub || !okRole {
				log.Println("❌ JWT missing or invalid 'sub' or 'role'")
				return fiber.ErrUnauthorized
			}

			role := int(roleFloat)
			log.Printf("✅ JWT role: %d, sub: %s", role, sub)
			// Store in context
			ctx.Locals("user_id", sub)
			ctx.Locals("user", sub)
			ctx.Locals("role", role)

			log.Println("🧪 Getting MAC and IP address...")
			ip, mac := get_mac_ip_address(ctx)
			log.Println("✅ IP and MAC fetched:", ip, mac)
			log.Printf("IP Address: %s", ip)
			log.Printf("MAC Address: %s", mac)
			ctx.Locals("ip_address", ip)
			ctx.Locals("mac_address", mac)
			ctx.Locals("encryption_key", os.Getenv("ENCRYPTION_KEY"))
			log.Println("🔐 Loaded encryption key:", os.Getenv("ENCRYPTION_KEY"))

			// Special access mapping
			pegawai := []int{1, 1337, 2, 3, 4001, 4002, 4003, 4004, 5001}

			// Allow all roles if 0
			if len(roles) == 1 && roles[0] == 0 {
				return ctx.Next()
			}

			for _, r := range roles {
				if role == r {
					return ctx.Next()
				}
				if r == 2 {
					for _, p := range pegawai {
						if role == p {
							return ctx.Next()
						}
					}
				}
			}

			log.Printf("❌ Forbidden: role %d not allowed", role)
			return fiber.ErrForbidden
		},

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Printf("❌ JWT error: %v", err)
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized access: " + err.Error(),
			})
		},

		SigningKey: jwtware.SigningKey{
			JWTAlg: jwt.SigningMethodHS512.Alg(),
			Key:    []byte(secret),
		},

		ContextKey: "jwt",
	})
}
