# go-vision-api

API de Extração de Dados de Documentos 

Esta API em Go utiliza a Google Cloud Vision API para extrair texto de diferentes tipos de documentos, como RG, CNH e passaporte.

Foram criados endpoints para cada tipo de documento suportado, permitindo a extração das informações a partir de imagens.

Endpoint /extract-rg: Extrai informações de RG a partir de uma imagem. 

Endpoint /extract-cnh: Extrai informações de CNH a partir de uma imagem. 

Endpoint /extract-passport: Extrai informações de passaporte a partir de uma imagem.

Exemplo do request: Método GET - http://localhost:8080/extract-rg?url=https://istoe.com.br/wp-content/uploads/2020/08/rg-digital.jpg
