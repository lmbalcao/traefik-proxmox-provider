# Doc/Code Alignment Report

- repo analisado: `traefik-proxmox-provider`
- ficheiros/documentacao inspecionados: `README.md`, `AGENTS.md`, `docs/README.md`, `docs/STATE.md`, `provider/provider.go`, `provider/provider_test.go`, `internal/service_test.go`, `proxmox-tagger.sh`
- evidencia principal encontrada: o codigo implementa `labelPrefix`, `pollInterval` minimo de `5s`, leitura de labels no campo notes/description e suporte testado para routers, middlewares, TLS, health checks, sticky cookies e backend scheme
- inconsistencias encontradas: o README anterior omitia `labelPrefix`, fixava uma versao de plugin nao revalidada e fazia afirmacoes demasiado largas sobre suporte total sem remeter para o que esta efetivamente coberto pelo codigo/testes
- correcoes aplicadas: `README.md` ajustado para refletir opcoes/config real e para evitar claims nao verificadas; criado este relatorio
- validacoes executadas: `bash -n scripts/validate-repo.sh proxmox-tagger.sh`; `python3 -m py_compile scripts/update-changelog.py`
- limitacoes / pontos nao validados: o ambiente nao disponibiliza `go`, por isso nao foi possivel correr `go test ./...`; nao foi estabelecida ligacao a um cluster Proxmox real nem a um Traefik real nesta auditoria
- resultado final: docs alinhadas
