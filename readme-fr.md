# Groupie Tracker - Absolut Music
Absolut Music est une application web développée dans le cadre du projet Groupie Tracker à l'école informatique Ynov.<br> L'application récupère des informations depuis une API pour permettre aux utilisateurs de suivre les concerts de leurs groupes de musique préférés.

## Prérequis

Avant de pouvoir utiliser cette application, vous devez avoir les éléments suivants :<br>

* Docker (version 23.0.1 ou ultérieure) (si vous souhaitez lancer l'application avec Docker)
* Go (version 1.19.6 ou ultérieure) (si vous souhaitez lancer l'application en utilisant ```go build```)

## Installation

1. Clonez ce dépôt de code sur votre ordinateur en utilisant la commande suivante :
``` bash
git clone https://github.com/TomJegou/Groupie-tracker.git
```

2. Accédez au répertoire de l'application en utilisant la commande suivante :
``` bash
cd Groupie-tracker
```

## Lancement avec Docker

1. Lancez l'application en utilisant la commande suivante :
``` bash
docker-compose up
```

2. Ouvrez votre navigateur web et accédez à l'URL suivante :
``` bash
http://localhost:80
```
3. Utilisez l'application pour rechercher des groupes de musique et suivre leurs concerts.

## Exécution avec Go

1. Compilez l'application en utilisant la commande suivante :
``` bash
go build -o bin/absolut-music
```

2. Lancez l'application en utilisant la commande suivante :
```bash
./bin/absolut-music
```

3. Ouvrez votre navigateur web et accédez à l'URL suivante :
``` bash
http://localhost:8080
```

4. Utilisez l'application pour rechercher des groupes de musique et suivre leurs concerts.