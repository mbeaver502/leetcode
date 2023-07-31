If you don't have composer installed on your local machine,
you'll have to run composer commands within the Docker container.

1. Identify the Docker container.
```bash
$ docker ps

CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS                NAMES
9924c84ea4e1   php-php   "docker-php-entrypoi…"   31 minutes ago   Up 31 minutes   0.0.0.0:80->80/tcp   php-apache
```

2. Enter the container in interactive mode and run bash.

```bash
$ docker exec -it CONTAINER_ID bash
```

3. Navigate to the desired directory and run your composer commands.

```bash
root@container:/your-directory$ composer require --dev phpunit/phpunit ^10
```

You can then run tests with PHPUnit in this way. For example:

```bash
root@container:/var/www/html$ cd two-sum
root@container:/var/www/html/two-sum$ ./vendor/bin/phpunit tests

PHPUnit 10.2.6 by Sebastian Bergmann and contributors.

Runtime:       PHP 8.2.8

.....                                                               5 / 5 (100%)

Time: 00:00.018, Memory: 6.00 MB

OK (5 tests, 14 assertions)

```