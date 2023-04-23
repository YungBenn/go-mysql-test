
## API Reference

#### Get All Mahasiswa
```http
  GET /mahasiswa
```


#### Get Mahasiswa Detail by ID

```http
  GET /mahasiswa/${id}
```

#### Insert New Mahasiswa

```http
  POST /mahasiswa/insert
```
- Contoh body raw json
```json
  {
      "nama": "Ruben",
      "usia": 21,
      "gender": 0,
      "jurusan": "Sistem Informasi",
      "hobi": "Nonton film",
  }
```

#### Update Mahasiswa by ID

```http
  PUT /mahasiswa/update/${id}
```
- Contoh body raw json
```json
  {
      "nama": "Ruben",
      "usia": 21,
      "gender": 0,
      "jurusan": "Sistem Informasi",
      "hobi": "Main game",
  }
```

#### Delete Mahasiswa by ID

```http
  DELETE /mahasiswa/delete/${id}
```
