# Shamelessly stolen from https://stackoverflow.com/a/69124635
FROM php:apache

RUN apt-get update && apt-get install -y \
    git \
    zip \
    curl \
    sudo \
    unzip \
    libzip-dev \
    libicu-dev \
    libbz2-dev \
    libpq-dev 

# 2. apache configs + document root
ENV APACHE_DOCUMENT_ROOT=/var/www/html/
RUN sed -ri -e 's!/var/www/html!${APACHE_DOCUMENT_ROOT}!g' /etc/apache2/sites-available/*.conf
RUN sed -ri -e 's!/var/www/!${APACHE_DOCUMENT_ROOT}!g' /etc/apache2/apache2.conf /etc/apache2/conf-available/*.conf

# 3. mod_rewrite for URL rewrite and mod_headers for .htaccess extra headers like Access-Control-Allow-Origin-
# RUN a2enmod rewrite headers
RUN a2enmod rewrite

# 4. start with base php config, then add extensions
# See https://github.com/mlocati/docker-php-extension-installer#supported-php-extensions
RUN mv "$PHP_INI_DIR/php.ini-development" "$PHP_INI_DIR/php.ini"
RUN docker-php-ext-install \
    bz2 \
    intl \
    # mysqli \
    # pgsql \
    # pdo \
    # pdo_mysql \
    # pdo_pgsql \
    # redis \
    zip 

RUN docker-php-ext-enable \
    bz2 \ 
    intl \
    # mysqli \
    # pgsql \
    # pdo \
    # pdo_mysql \
    # pdo_pgsql \
    # redis \
    zip 

# RUN pecl install -o -f redis \
#     &&  rm -rf /tmp/pear \
#     &&  docker-php-ext-enable redis

# 5. composer
COPY --from=composer /usr/bin/composer /usr/bin/composer