# Compilando :gear:

Para compilar para outro sistema, você precisara ter instalado:

- [***Golang***](https://golang.org)
- O [***Source Code***](https://github.com/Henriquetied472/tabshell-cli)

---

#### Primeiro passo :gear:

Clone este repositório para sua máquina local e entre nele usando:

```shell
git clone https://github.com/Henriquetied472/tabshell-cli.git
cd tabshell-cli
```

#### Segundo passo :gear:

Compile o código para a plataforma que quiser usando o `GOOS` para definir o Sistema Operacional alvo e o `GOARCH` definir a Arquitetura alvo:

```shell
GOOS=<OS> GOARCH=<ARCH> go build main.go
mv main tabshell
```

**OBS: Troque o `<OS>` pelo Sistema Operacional e `<ARCH>` pela Arquitetura**

---

Para ver a lista de Sistemas Operacionais e Arquiteturas, use este comando:

```shell
go tool dist list
<OS>/<ARCH>
```
