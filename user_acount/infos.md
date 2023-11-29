# Mise en place du systeme de compte utilisateur

## Le hashing :

 - Utilisation : le SHA-2 dans une de ses implémentations : SHA-256, SHA-512, SHA-384, SHA-224.
 - Mise en place : a la création du compte, conection, modification de mot de passe ou de tout autre paramétre

Solution : un seul programe qui prend en paramétre une string et retourne son hash

## Intéraction avec la basse de données :

 - Création : on doit pouvoir créer de nouveaux comptes.

 Solution : un seul programe qui prend en compte la basse de donnée et les éléments a créer et crée les données dans la table.

 - Récupération : les comptes néscecite de récupérer des éléments.

 Solution : un seul programe qui prend en compte les information a récupérer (basse de donnée visé, élements ciblé.....) et retourne les éléments qui se trouve dedant.

  - Modification : les comptes peuvent étre modifié.

  Solution : un seul programe qui prend en compte la basse de donnée, l'element a modifier et l'élément qui va le remplacer.

## Sécurité :

 - le front sécurisé.
 - les intéraction front/back sécurisé.
 - sécurité des comptes (hashing des mots de passe).
 - 