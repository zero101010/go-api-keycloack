# go-api-keycloack

- Iremos criar 2 apis para se comunicar entre elas usando keycloack e golang


## Fincionamento keycloack
- Resource Owner (É o dono do recurso, no nosso caso o cliente que tenta acessar a aplicação)
- Client (É o sistema que quer ter acesso aquele recurso)
- Resource Server( O sistema que o recurso está)
- Authorization server ( O recurso que valida se o client pode ter acesso ao Resource Server)

Usaremos o keycloack como docker inicialmente usando o seguinte comando:
```
docker run -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin quay.io/keycloak/keycloak:12.0.4
```
### Estrutura keycloack
- Criar realm que é como se fosse uma workspace de trabalho, pode ser usado como camada de serviço entre aplicações para termos mais de um sistema autenticando em apenas um keycloack
- Criar um client que é como se fosse o contexto da sua aplicação
- Criar um User para tentar acessar a nossa aplicação

### Aplicação
- Foi criado um código simples em Go para usar o servidor de autenticação do keycloack como forma de autenticar

