# watermelon-db-api

Backend do projeto **WatermelonDB**: uma plataforma colaborativa para o compartilhamento de índices de qualidade de água.

## Instalação com docker compose
Use `docker compose up --build -d` para criar os contâineres do projeto, criando o banco de dados e populando com exemplos. Por padrão, a API aguardará por requisições na porta `8080`. Para mudar isso, altere o seguinte campo no arquivo `compose.yaml`:
```
ports:
      # Altere para a porta que você deseja expor
      - "8080:8080"
```

Por padrão, a API aceitará apenas os métodos POST, PATCH e DELETE vindos de um *host* específico. Por padrão, esse host é o `http://localhost:5000`. Para usar outro nome (como `http://127.0.0.1` ou `https://exemplo.com.br`), altere o seguinte campo no arquivo `compose.yaml`:
```
environment:
  # Altere para o endereço do front end
  REQUEST_HOST: http://localhost:5000
```

Encerre a aplicação com `docker compose down`.

## Instalação de forma nativa
- Altere o endereço da conexão com o banco de dados no arquivo `banco/banco.go` para usar o local do seu servidor MySQL/MariaDB:
  ```
  func Banco() *gorm.DB {
    // Siga o formato "usuario:senha@tcp(endereço:porta)/watermelon_db".
    db, err := gorm.Open(mysql.Open("watermelon:watermelon@tcp(db:3306)/watermelon_db"), &gorm.Config{
  ```

- Altere o endereço do front-end no arquivo `main.go` para que a API possa aceitar as requisições:
  ```
  func cors(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      // Remova 'os.Getenv("REQUEST_HOST")' e subsitua pelo seu endereço (como "http://localhost:8000" ou "https://exemplo.com")
	  w.Header().Set("Access-Control-Allow-Origin", os.Getenv("REQUEST_HOST"))
  ```
  Remova também a importação no mesmo arquivo `main.go`:
  ```
  import (
  	"api/auth"
  	"api/crud"
  	"os"  // <------ REMOVA OU COMENTE ESTA LINHA!
  
  	"log"
  	"net/http"
  )
  ```

### Dependências
- `go` 1.25.3
- `mysql` ou `mariadb`
- Instale as bibliotecas com `go mod tidy`

### Execução
1) Crie e popule o banco de dados. Use:
   ```
   mysql -u root -p < ./sql/01-watermelondb.sql
   mysql -u root -p < ./sql/02-cidades.sql
   mysql -u root -p < ./sql/03-popular.sql
   ```
     
2) Compile o projeto com `go build .` e execute com `./api`. Alternativamente, você pode executar o projeto com `go run .` (não gera um arquivo executável).
