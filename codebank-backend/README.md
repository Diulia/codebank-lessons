<h1>Como rodar a aplicação: </h1>
<p>
    <ol>
        <li>Primeiramente, abra a pasta "codebank-backend"</li>
        <li>Rode primeiramente o docker-compose ps para verificar se já existem os containers e se eles estão rodando</li>
        <li>Se estiverem com o status exit ou qualquer outro, rode um docker-compose down e em seguida docker-compose up -d</li> 
        <li>Depois que estiver rodando certo é só acessar a bash do aplicativo com o comando docker exec -it appbank bash </li>
        <li> rode o comando go run main.go dentro do bash para rodar a aplicação</li>
    </ol>

</p>