<div align="right" >

##### [:angola: PT](../README.md) &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp; :us: EN

</div>

<div align="center">

<img alt="NLW Spacetime" src="/docs/images/logo.png" width="100px"  />

</p>

 [Project](#-project)&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[Technologies](#rocket-technologies)&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[Run](#gear-run)&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[Layout](#nail_care-layout)&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[License](#-license)
</div>

# NLW Spacetime • Ignite Trail

![Cover](/docs/images/cover.png)

## 💻 Project

Application to record your memories, where the user can add texts and photos of important events in his life to a timeline, organized by month and year.

</br>

## :rocket: Technologies

This project was developed with the following technologies:

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

## :gear: Run

Clone the project to a local directory:</p>

```bash
git clone https://github.com/EdlanioJ/nlw-spacetime.git
```

Create a new [Oauth Github](https://github.com/settings/apps) project and copy yours credentials to `.env` file.

### Server

```bash
cd server
```

Create a `.env` file with [env.example](../server/.env.example) template found in the server folder.

Create a free account on [UploadCare](https://uploadcare.com/), create a project, create the credentials and copy `.env` file.

install [Golang](https://go.dev/) e:

```bash
  go mod download

  go run main.go
```

or with [Docker](https://www.docker.com/):

```bash
docker compose up
```

### Web

```bash
cd web

npm install
```

Create a `.env` file with [env.example](../web/.env.example) template found in the web folder.

```bash
npm run dev
```

### Mobile

```bash
cd web

npm install
```

Create a `.env` file with [env.example](../mobile/.env.example) template found in the mobile folder.

```bash
npx expo start
```

## :nail_care: Layout

You may view the layout of this project through this [link :link:](https://www.figma.com/file/dCOA0zAWoSge4yRiyzVXn4/C%C3%A1psula-do-tempo-%E2%80%A2-Trilha-Ignite-(Community)?type=design&node-id=206-157&t=PKFsdAi7pDIdgYSK-0). You must have an account at [Figma](https://figma.com) to access it.

## 📝 License

This project is under the MIT license. See the [LICENSE](LICENSE) file for more details.

---

Made with :heart: por [Edlâneo Manuel](https://github.com/EdlanioJ) :wave:
