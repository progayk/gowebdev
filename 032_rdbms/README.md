1. install mysql
1. install workbench
1. create mysql db on aws
1. connect workbench to rds mysql db
1. https://www.youtube.com/watch?v=k68Y-XYapEI


**resources & credits**

* https://github.com/golang/go/wiki/SQLInterface
* https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-18-04

#Install MySQL on local machine
##Introduction

update your package index, install the mysql-server package, and then run the included security script.

```$xslt
sudo apt update
sudo apt install mysql-server
sudo mysql_secure_installation
```

##Step 1 - Install MySQL
To install it, update the package index on your server with apt:
```$xslt
sudo apt update
```

Then install the default package:
```$xslt
sudo apt install mysql-server
```

##Step 2 - Configuring MySQL
Run the security script:
```$xslt
sudo mysql_secure_installation
```
This will prompt some questions. Create a **root** user.

To initialize the MySQL data directory, you would use mysql_install_db for versions before 5.7.6, and mysqld --initialize for 5.7.6 and later. However, if you installed MySQL from the Debian distribution, as described in Step 1, the data directory was initialized automatically; you don't have to do anything. If you try running the command anyway, you'll see the following error:

```$xslt
// Output:
mysqld: Can't create directory '/var/lib/mysql/' (Errcode: 17 - File exists)
. . .
2018-04-23T13:48:00.572066Z 0 [ERROR] Aborting
```
Note that even though you’ve set a password for the root MySQL user, this user is not configured to authenticate with a password when connecting to the MySQL shell. If you’d like, you can adjust this setting by following Step 3.
##Step 3 - (Optional) Adjusting User Authentication and Privileges
In Ubuntu systems running MySQL 5.7 (and later versions), the root MySQL user is set to authenticate using the auth_socket plugin by default rather than with a password. This allows for some greater security and usability in many cases, but it can also complicate things when you need to allow an external program (e.g., phpMyAdmin) to access the user.

In order to use a password to connect to MySQL as root, you will need to switch its authentication method from auth_socket to mysql_native_password. To do this, open up the MySQL prompt from your terminal:
```$xslt
sudo mysql
```
Next, check which authentication method each of your MySQL user accounts use with the following command:
```mysql
SELECT user,authentication_string,plugin,host FROM mysql.user;
```

```$xslt
Output:
+------------------+-------------------------------------------+-----------------------+-----------+
| user             | authentication_string                     | plugin                | host      |
+------------------+-------------------------------------------+-----------------------+-----------+
| root             |                                           | auth_socket           | localhost |
| mysql.session    | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE | mysql_native_password | localhost |
| mysql.sys        | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE | mysql_native_password | localhost |
| debian-sys-maint | *CC744277A401A7D25BE1CA89AFF17BF607F876FF | mysql_native_password | localhost |
+------------------+-------------------------------------------+-----------------------+-----------+
4 rows in set (0.00 sec)
```

In this example, you can see that the **root** user does in fact authenticate using the `auth_socket` plugin. To configure the **root** account to authenticate with a password, run the following `ALTER USER` command. Be sure to change `password` to a strong password of your choosing, and note that this command will change the **root** password you set in Step 2:

```mysql
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
```
Then, run `FLUSH PRIVILEGES` which tells the server to reload the grant tables and put your new changes into effect:

```mysql
FLUSH PRIVILEGES;
```
Check the authentication methods employed by each of your users again to confirm that root no longer authenticates using the auth_socket plugin:
```mysql
SELECT user,authentication_string,plugin,host FROM mysql.user;
```
```$xslt
Output
+------------------+-------------------------------------------+-----------------------+-----------+
| user             | authentication_string                     | plugin                | host      |
+------------------+-------------------------------------------+-----------------------+-----------+
| root             | *3636DACC8616D997782ADD0839F92C1571D6D78F | mysql_native_password | localhost |
| mysql.session    | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE | mysql_native_password | localhost |
| mysql.sys        | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE | mysql_native_password | localhost |
| debian-sys-maint | *CC744277A401A7D25BE1CA89AFF17BF607F876FF | mysql_native_password | localhost |
+------------------+-------------------------------------------+-----------------------+-----------+
4 rows in set (0.00 sec)
```

You can see in this example output that the root MySQL user now authenticates using a password. Once you confirm this on your own server, you can exit the MySQL shell:

```mysql
exit
```
Alternatively, some may find that it better suits their workflow to connect to MySQL with a dedicated user. To create such a user, open up the MySQL shell once again:
```$xslt
sudo mysql
```

:::
**Note:** If you have password authentication enabled for root, as described in the preceding paragraphs, you will need to use a different command to access the MySQL shell. The following will run your MySQL client with regular user privileges, and you will only gain administrator privileges within the database by authenticating:
```$xslt
mysql -u root -p
```
:::

From there, create a new user and give it a strong password:
```mysql
CREATE USER 'sammy'@'localhost' IDENTIFIED BY 'password';
```

Then, grant your new user the appropriate privileges. 
```mysql
GRANT ALL PRIVILEGES ON *.* TO 'sammy'@'localhost' WITH GRANT OPTION;
```

Because you created a new user, instead of modifying an existing one, `FLUSH PRIVILEGES` is unnecessary here.

Finally, let's test the MySQL installation.

##Step 4 — Testing MySQL

Regardless of how you installed it, MySQL should have started running automatically. To test this, check its status.
```$xslt
systemctl status mysql.service
```
If MySQL isn't running, you can start it with `sudo systemctl start mysql`.

For an additional check, you can try connecting to the database using the mysqladmin tool, which is a client that lets you run administrative commands. For example, this command says to connect to MySQL as root (-u root), prompt for a password (-p), and return the version.
```$xslt
sudo mysqladmin -p -u root version
```

# MySQL - Basic Usage
 
## Creating Database

**Show existing databases:**
```mysql
SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.00 sec)
```

**Creating a new database:**
```mysql
CREATE DATABASE pets;
Query OK, 1 row affected (0.01 sec)
```
Check if the database has been created:
```mysql
SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| pets               |
| sys                |
+--------------------+
5 rows in set (0.00 sec)
```

**Creating a table inside a database**.  First, pick the database in which you want to create the table with a USE statement:
```mysql
USE pets
Database changed
```
The USE statement tells MySQL to use pets as the default database for subsequent statements. Next, create a table with a CREATE TABLE statement:
```mysql
CREATE TABLE cats
(
  id              INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
  name            VARCHAR(150) NOT NULL,                # Name of the cat
  owner           VARCHAR(150) NOT NULL,                # Owner of the cat
  birth           DATE NOT NULL,                        # Birthday of the cat
  PRIMARY KEY     (id)                                  # Make the id the primary key
);
```

Data types you can use in each column are explained in Data Types. Primary Key Optimization explains the concept of a primary key. What follows a # on each line is a comment, which is ignored by the mysql client; see [Comment Syntax](https://dev.mysql.com/doc/refman/8.0/en/comments.html) for other comment styles.
```mysql
mysql> SHOW TABLES;
+----------------+
| Tables_in_pets |
+----------------+
| cats           |
+----------------+
1 row in set (0.00 sec)
```

DESCRIBE shows information on all columns of a table:
```mysql
DESCRIBE cats;
+-------+------------------+------+-----+---------+----------------+
| Field | Type             | Null | Key | Default | Extra          |
+-------+------------------+------+-----+---------+----------------+
| id    | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| name  | varchar(150)     | NO   |     | NULL    |                |
| owner | varchar(150)     | NO   |     | NULL    |                |
| birth | date             | NO   |     | NULL    |                |
+-------+------------------+------+-----+---------+----------------+
4 rows in set (0.00 sec)
```

**Adding records into a table**.  Use, for example, an INSERT...VALUES statement:

```mysql
INSERT INTO cats ( name, owner, birth) VALUES
  ( 'Sandy', 'Lennon', '2015-01-03' ),
  ( 'Cookie', 'Casey', '2013-11-13' ),
  ( 'Charlie', 'River', '2016-05-21' );
```

See Literal Values for how to write string, date, and other kinds of literals in MySQL.

**Retrieving records from a table**.  Use a SELECT statement, and “*” to match all columns:

```mysql
mysql> SELECT * FROM cats;
+----+---------+--------+------------+
| id | name    | owner  | birth      |
+----+---------+--------+------------+
|  1 | Sandy   | Lennon | 2015-01-03 |
|  2 | Cookie  | Casey  | 2013-11-13 |
|  3 | Charlie | River  | 2016-05-21 |
+----+---------+--------+------------+
3 rows in set (0.00 sec)
```

To select specific columns and rows by a certain condition using the WHERE clause:

```mysql
SELECT name FROM cats WHERE owner = 'Casey';
+--------+
| name   |
+--------+
| Cookie |
+--------+
1 row in set (0.00 sec)
```
**Deleting a record from a table**.  Use a DELETE statement to delete a record from a table, specifying the criterion for deletion with the WHERE clause:

```mysql
DELETE FROM cats WHERE name='Cookie';
Query OK, 1 row affected (0.05 sec)

mysql> SELECT * FROM cats;
+----+---------+--------+------------+
| id | name    | owner  | birth      |
+----+---------+--------+------------+
|  1 | Sandy   | Lennon | 2015-01-03 |
|  3 | Charlie | River  | 2016-05-21 |
+----+---------+--------+------------+
2 rows in set (0.00 sec)
```

**Adding or deleting a column from a table**.  Use an ALTER TABLE...ADD statement to add a column. You can use, for example, an AFTER clause to specify the location of the new column:

```mysql
ALTER TABLE cats ADD gender CHAR(1) AFTER name;
Query OK, 0 rows affected (0.24 sec)
Records: 0  Duplicates: 0  Warnings: 0
```

Use DESCRIBE to check the result:

```mysql
DESCRIBE cats;
+--------+------------------+------+-----+---------+----------------+
| Field  | Type             | Null | Key | Default | Extra          |
+--------+------------------+------+-----+---------+----------------+
| id     | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| name   | varchar(150)     | NO   |     | NULL    |                |
| gender | char(1)          | YES  |     | NULL    |                |
| owner  | varchar(150)     | NO   |     | NULL    |                |
| birth  | date             | NO   |     | NULL    |                |
+--------+------------------+------+-----+---------+----------------+
5 rows in set (0.00 sec)
```

`SHOW CREATE TABLE shows a CREATE TABLE statement`, which provides even more details on the table:

```mysql
SHOW CREATE TABLE cats\G
*************************** 1. row ***************************
       Table: cats
Create Table: CREATE TABLE `cats` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(150) NOT NULL,
  `gender` char(1) DEFAULT NULL,
  `owner` varchar(150) NOT NULL,
  `birth` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1
1 row in set (0.00 sec)
```

Use `ALTER TABLE...DROP` to delete a column:

```mysql
ALTER TABLE cats DROP gender;
Query OK, 0 rows affected (0.19 sec)
Records: 0  Duplicates: 0  Warnings: 0

DESCRIBE cats;
+-------+------------------+------+-----+---------+----------------+
| Field | Type             | Null | Key | Default | Extra          |
+-------+------------------+------+-----+---------+----------------+
| id    | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| name  | varchar(150)     | NO   |     | NULL    |                |
| owner | varchar(150)     | NO   |     | NULL    |                |
| birth | date             | NO   |     | NULL    |                |
+-------+------------------+------+-----+---------+----------------+
4 rows in set (0.00 sec)
```

