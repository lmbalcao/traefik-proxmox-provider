package internal

import (
	"strings"
	"log"
)

type ParsedConfig struct {
	Description string `json:"description,omitempty"`
}

type ParsedAgentInterfaces struct {
     Data []struct {
	     Name        string `json:"name"`
             IPAddresses []IP `json:"ip-addresses"`
     } `json:"data"`
}

type NodeStatus struct {
	Node string `json:"node"`
}

type VirtualMachine struct {
	VMID   uint64 `json:"vmid"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Container struct {
	VMID   uint64 `json:"vmid"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Version struct {
	Release string `json:"release"`
}

type Service struct {
	ID     uint64
	Name   string
	IPs    []IP
	Config map[string]string
}

type IP struct {
	Address     string `json:"ip-address,omitempty"`
	AddressType string `json:"ip-address-type,omitempty"`
	Prefix      uint64 `json:"prefix,omitempty"`
}

func NewService(id uint64, name string, config map[string]string) Service {
	return Service{ID: id, Name: name, Config: config, IPs: make([]IP, 0)}
}

func (pc *ParsedConfig) GetTraefikMap(labelPrefix string) map[string]string {
        const separator = "="

        prefix := strings.ToLower(strings.TrimSpace(labelPrefix))
        if prefix == "" {
                prefix = "traefik."
        }

        // Mantém a lógica antiga:
        // converte labels separadas por espaço em labels separadas por newline,
        // mas usando o prefixo configurado.
        normalized := strings.ReplaceAll(pc.Description, " "+prefix, "\n"+prefix)

        m := make(map[string]string)
        lines := strings.Split(normalized, "\n")
        for _, line := range lines {
                key, value, found := strings.Cut(line, separator)
                if !found {
                        continue
                }

                key = strings.Trim(key, "\" ")
                value = strings.Trim(value, "\" ")

                lowerKey := strings.ToLower(key)
                if strings.HasPrefix(lowerKey, prefix) {
                        suffix := strings.TrimPrefix(lowerKey, prefix)
                        normalizedKey := "traefik." + suffix
                        m[normalizedKey] = value
                }
        }

        return m
}

func (pai *ParsedAgentInterfaces) GetIPs() []IP {
        primary := make([]IP, 0)
        secondary := make([]IP, 0)

        for _, r := range pai.Data {
                log.Printf("DEBUG iface=%s ips=%+v", r.Name, r.IPAddresses)

                isPrimary := r.Name == "eth0" || r.Name == "ens18"

                for _, ip := range r.IPAddresses {
                        if isPrimary {
                                primary = append(primary, ip)
                        } else {
                                secondary = append(secondary, ip)
                        }
                }
        }

        log.Printf("DEBUG primary=%+v secondary=%+v", primary, secondary)

        return append(primary, secondary...)
}
