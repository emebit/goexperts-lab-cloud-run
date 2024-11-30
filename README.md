# Goexperts-Lab-Cloud-Run

Um sistema de consulta de temperatura por CEP.

## Como usar

### Pré-requisitos

- Go (1.22 ou maior) instalado
- Docker (opcional)

### Executando localmente

1. Clone o repositório:

    ```bash
    git clone git@github.com:emebit/goexperts-lab-cloud-run.git
    cd goexperts-lab-cloud-run

2. Compile e execute o programa:

    go run cmd/main.go 
    Server is running 8080 

3. Fazer a requisição que está no arquivo http/test.http. O resultado será parecido com o seguinte:
   
    HTTP/1.1 200 OK 
    Content-Type: application/json 
    Date: Wed, 06 Nov 2024 00:14:31 GMT 
    Content-Length: 59 

    {
    "temp_C": 20.2,
    "temp_F": 68.36,
    "temp_K": 293.34999999999997
    } 

### Link cloud

https://goexperts-lab-cloud-run-272371299955.us-central1.run.app/cep/82015680

### Executando com Docker

Você pode construir uma imagem docker e executar a aplicação.

1. Build da imagem:

    docker build -t goexperts-lab-cloud-run .

2. Execute a imagem:

    docker run -p 8080:8080  goexperts-lab-cloud-run 

3. Fazer a requisição que está no arquivo http/test.http. O resultado será o similar ao seguinte:
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Sat, 30 Nov 2024 21:42:25 GMT
    Content-Length: 71
    Connection: close
    
    {
      "temp_C": 22.2,
      "temp_F": 71.96000000000001,
      "temp_K": 295.34999999999997
    }
