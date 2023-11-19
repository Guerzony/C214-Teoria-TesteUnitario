O Projeto consiste em gerenciador de alunos

# Equipe

- Vinicius Ribas
- Pietro de Souza Cardoso
- Gabriel Zordan
- Gabriel Ribeiro

# Framework

> Projeto feito utilizando o Go lang, framework utilizado na disciplina de C214, do Professor Christopher Lima.

# Instalação

## Go 

### Linux

#### 1. Download: 
```bash
wget https://golang.org/dl/go1.17.linux-amd64.tar.gz
```
#### 2. Extrair e Instalar:
```bash
sudo tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz
```
#### 3. Configurar as Variaveis de Ambiente:
Adicione as seguintes linhas no final do seu arquivo ~/.bashrc ou ~/.zshrc (ou equivalente, dependendo do seu shell):
```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
Em seguida, recarregru o arquivo de perfil do shell:
```bash
source ~/.bashrc
```

### macOS:

#### 1. Download: 
```bash
 brew install go
```
#### 2. Configurar as Variaveis de Ambiente:
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
Em seguida, recarregue o arquivo de perfil do shell
```bash
source ~/.zshrc
```

### Windows

#### 1. Download:
Baixe o instalador MSI do <a href="https://go.dev/dl/" >site oficial do Go</a> e execute o instalador.


### Verificar a instalação 

Para verificar se o Go foi instalado corretamente abra um terminal e execute o comando:
```bash
go version
```
Você sera capaz de ver a versão do GO instalada.

Se você entende um pouco de inglês, e quer ler a documentação oficial de instalação do Go lang basta <a href="https://go.dev/doc">clicar aqui</a>

### Docker

Para uso de banco de dados vamos usar o docker, para tal devera ser baixado e instalado de acordo com as instruções do site <a href="https://www.docker.com/get-started/">Docker</a>. 

### Visual Studio Code

- Baixe e instale <a href="https://code.visualstudio.com/download">Visual Studio Code</a>
- Inicie o Visual Studio Code e vá para 'Extencions'. Procure por Go lang e baixe a sua extenção.

# Execução



### Instalação das depedencias

Para instalação das dependencias do projeto basta rodar
```bash
go mod tidy
```

# Rodando o codigo

Para rodar o codigo primeiro precisamos subir o container docker, para o mesmo basta seguir o comando:
```bash
docker-compose up
```

Após o container estiver ativo basta rodar o comando:

```bash
go run main.go
```

### Comandos de teste

Para executar os testes basta rodar:

```bash
go test ./...
```
Para um teste mais detalhado basta rodar esse comando:
```bash
go test -v ./...
```
Caso queira verificar a cobertura dos testes rode:
```bash
go test -cover ./...
```

### Utilizando a aplicação

A aplicação possui uma interface bem intuitiva, com a entrada de todas as funcionalidades do codigo como: ExibeAlunos, NovoAluno, Buscar, Delete, Editar, BuscaCPF, BuscaRG.
