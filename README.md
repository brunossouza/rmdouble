# rmdouble

A command line tool to identify and remove duplicated files based on MD5 hash comparison.

⚠️ **ATENÇÃO: Este é um código de estudo e não deve ser usado em produção!**

## Descrição

Este programa recursivamente percorre diretórios, calcula o hash MD5 de cada arquivo e identifica duplicatas. Opcionalmente, pode contar ou deletar os arquivos duplicados encontrados.

## Uso

```bash
go run main.go [flags]
```

### Flags disponíveis:

- `-p <path>`: Caminho do diretório a ser verificado (padrão: `./`)
- `-r`: Modo recursivo - percorre subdiretórios
- `-d`: Deleta arquivos duplicados encontrados
- `-c`: Conta o número de arquivos duplicados
- `-v`: Modo verbose - exibe informações detalhadas

### Exemplos:

```bash
# Verificar diretório atual com saída detalhada
go run main.go -v

# Verificar recursivamente um diretório específico
go run main.go -p /caminho/para/diretorio -r -v

# Contar duplicados sem deletar
go run main.go -p /caminho/para/diretorio -r -c -v

# CUIDADO: Deletar duplicados (use com extrema cautela!)
go run main.go -p /caminho/para/diretorio -r -d -v
```

## ⚠️ Avisos Importantes

- **NÃO USE EM PRODUÇÃO**: Este código foi desenvolvido para fins educacionais
- **FAÇA BACKUP**: Sempre faça backup dos seus arquivos antes de usar a flag `-d`
- **TESTE PRIMEIRO**: Use as flags `-v` e `-c` para verificar antes de deletar
- **SEM CONFIRMAÇÃO**: A flag `-d` deleta arquivos sem pedir confirmação

## Limitações

- Usa MD5 para comparação (não é seguro para uso criptográfico)
- Não verifica colisões de hash (embora sejam raras)
- Interface de linha de comando básica
- Sem mecanismo de desfazer operações de deleção

## Compilação

```bash
go build -o rmdouble main.go
```