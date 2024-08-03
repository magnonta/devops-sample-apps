# Use a imagem oficial do PHP
FROM php:8.0-apache

# Defina o diretório de trabalho dentro do container
WORKDIR /var/www/html

# Copie o código PHP para o container
COPY . .

# Expose port 80 to the outside world
EXPOSE 80

# Defina uma variável de ambiente padrão
ENV APP_ENV=prod

# Comando padrão para iniciar o servidor Apache com PHP
CMD ["apache2-foreground"]
