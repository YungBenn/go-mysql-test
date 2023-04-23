# go-mysql-test

Internship Technical Test

## Run Locally

Clone the project

```bash
  git clone https://github.com/YungBenn/go-mysql-test.git
```

Go to the project directory

```bash
  cd go-mysql-test
```

Install dependencies

```bash
  go get
```

Create a `.env` file in the root directory of the project and see `.env.example` for an example

-   `PORT`

-   `DB_URL`

Start the server

```bash
  go run cmd/main.go
```

## API Reference

#### Get All Mahasiswa

```http
  https://www.google.com/
```

#### Get Mahasiswa Detail by ID

```http
  GET /mahasiswa/${id}
```

#### Insert New Mahasiswa

```http
  POST /mahasiswa/insert
```

-   Contoh body raw json

```json
{
    "nama": "Ruben",
    "usia": 21,
    "gender": 0,
    "jurusan": "Sistem Informasi",
    "hobi": "Nonton film"
}
```

#### Update Mahasiswa by ID

```http
  PUT /mahasiswa/update/${id}
```

-   Contoh body raw json

```json
{
    "nama": "Ruben",
    "usia": 21,
    "gender": 0,
    "jurusan": "Sistem Informasi",
    "hobi": "Main game"
}
```

#### Delete Mahasiswa by ID

```http
  DELETE /mahasiswa/delete/${id}
```
