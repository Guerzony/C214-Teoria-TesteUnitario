O Projeto consiste em gerenciador de alunos

# Equipe

- Vinicius Ribas
- Pietro de Souza Cardoso
- Gabriel Zordan
- Gabriel Ribeiro

# Framework

> Projeto feito utilizando o Go lang, framework utilizado na disciplina de C214, do Professor Christopher Lima.

# Instalação

Se você entende um pouco de inglês, e quer ler a documentação oficial de instalação do Go lang basta <a href="https://go.dev/doc/install"/>clicar aqui</a>

### Visual Studio Code

- Baixe e instale <a href="https://code.visualstudio.com/download">Visual Studio Code</a>
- Inicie o Visual Studio Code e vá para 'Extencions'. Procure por Go lang e baixe a sua extenção.


### Instalação das depedencias

### Arquivo go.mod

No diretorio raiz do projeto execute o comando

```bash
go mod init nomeDoProjeto
```
isso criara o arquivo `go.mod` que gerencia as dependências do projeto.


### Comandos de teste

O comando abaixo roda todos os arquivos teste do projeto.

```bash
Go test ./...
```

### Utilizando a aplicação

A aplicação possui uma interface bem intuitiva, com a entrada de todas as funcionalidades do codigo como: ExibeAlunos, NovoAluno, Buscar, Delete, Editar, BuscaCPF, BuscaRG.
