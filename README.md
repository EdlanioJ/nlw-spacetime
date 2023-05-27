
# <img alt="NLW Spacetime" src="/docs/images/logo.png" style="display: block; margin: 0 auto" width="100px"  />

</p>

<center>

 [Projeto](#-projeto)&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[Tecnologias](#rocket-tecnologias)&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[Executar](#gear-executar)&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[Layout](#nail_care-layout)&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[Licença](#-licença)
</center>

# NLW Spacetime • Trilha Ignite

![Cover](/docs/images/cover.png)

## 💻 Projeto

Aplicação de recordação de memórias, onde o usuário poderá adicionar à uma timeline textos e fotos de acontecimentos marcantes da sua vida, organizados por mês e ano.

</br>

## :rocket: Tecnologias

Esse projeto foi desenvolvido com as seguintes tecnologias:

- [NodeJS](https://nodejs.org/en/)
- [Typescript](https://www.typescriptlang.org/)
- [Golang](https://go.dev/)

---

- Server
  - [x] [Fiber](https://docs.gofiber.io/)
  - [x] [sqlx](https://github.com/jmoiron/sqlx)
  - [x] [UploadCare](https://uploadcare.com/)
  - [x] [PostgresSQL](https://www.postgresql.org/)
  - [x] [Docker](https://www.docker.com/)
- Web
  - [x] [ReactJS](https://reactjs.org/)
  - [x] [Next.JS](https://nextjs.org/)
  - [x] [TailwindCSS](https://tailwindcss.com/)

- Mobile
  - [x] [React Native](https://reactnative.dev/)
  - [x] [Expo](https://expo.dev/)
  - [x] [NativeWind](https://www.nativewind.dev/)

## :gear: Executar

Clonar o projeto para um diretório local:</p>

```bash
git clone https://github.com/EdlanioJ/nlw-spacetime.git
```

Crie um novo projeto oauth no [Github](https://github.com/settings/apps) e copie as credencias para os ficheiros `.env`

### Server

```bash
cd server
```

Criar um ficheiro `.env` com o template [.env.example](/server/.env.example) encontrado na pasta server.

Criar uma conta gratis no [UploadCare](https://uploadcare.com/), crie um novo projeto e configure e copie as credenciais para o ficheiro `.env`

instalar [Golang](https://go.dev/) e:

```bash
  go run main.go
```

ou Com [Docker](https://www.docker.com/):

```bash
docker compose up
```

### Web

```bash
cd web

npm install
```

Criar um ficheiro `.env` com o template [.env.example](/web/.env.example) encontrado na pasta web.

```bash
npm run dev
```

### Mobile

```bash
cd web

npm install
```

Criar um ficheiro `.env` com o template [.env.example](/mobile/.env.example) encontrado na pasta mobile.

```bash
npx expo start
```

## :nail_care: Layout

Você pode visualizar o layout do projeto através do desse [link :link:](https://www.figma.com/file/dCOA0zAWoSge4yRiyzVXn4/C%C3%A1psula-do-tempo-%E2%80%A2-Trilha-Ignite-(Community)?type=design&node-id=206-157&t=PKFsdAi7pDIdgYSK-0). É necessário ter conta no [Figma](https://figma.com) para acessá-lo.

## 📝 Licença

Esse projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

Feito com :heart: por [Edlâneo Manuel](https://github.com/EdlanioJ) :wave:
