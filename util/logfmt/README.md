# logfmt

`logfmt` permet d'encoder des couples de clé/valeur au format ["logfmt"](https://www.brandur.org/logfmt).

**Exemples:**

| couple clé/valeur                             | résultat                                     |
|-----------------------------------------------|----------------------------------------------|
| nil, nil                                      | @nil=@nil                                    |
| age, 53                                       | age=53                                       |
| "a\tb\nc", "def"                              | abc="def"                                    |
| []byte("lsm"), "ceci est un message"          | []byte{0x6c,0x73,0x6d}="ceci est un message" |
| "", 789.456                                   | @key=789.456                                 |
| "jour", 24, "mois", "décembre", "année", 2019 | jour=24 mois="décembre" année=2019           |
| "la valeur est manquante"                     | lavaleurestmanquante=@nil                    |
| "message", "Joyeuses\tfêtes\n"                | message="Joyeuses\tfêtes\n"                  |

**Remarques:**

- `@key` remplace les clés qui ne sont pas valides.
- `@nil` correspond à une clé ou une valeur valant `nil`.

---

Copyright (c) 2020 Institut National de l'Audiovisuel