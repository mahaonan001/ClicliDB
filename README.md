      +---------------------+       purchases      +---------------------+
      |      customer       |----------------------|       ticket        |
      +---------------------+                      +---------------------+
      | c_ID (PK)           |<--------------------<| ticket_ID (PK)      |
      | name                |                      | seat_num            |
      | phone               |                      | c_ID (FK)           |
      +---------------------+                      | schedule_ID (FK)    |
                                                   +---------------------+
                                                           |
                                                           |
                                                           |
                                                           v
                                                   +---------------------+
                                                   |      schedule       |
                                                   +---------------------+
                                                   | schedule_ID (PK)    |
                                                   | date                |
                                                   | time                |
                                                   | price               |
                                                   | number              |
                                                   | movie_ID (FK)       |
                                                   | hall_ID (FK)        |
                                                   +---------------------+
                                                          /|\
                                                           |
                                                   screens  |
                                                           |
                                                           |
                                                          \|/
                                                   +---------------------+
                                                   |       movie         |
                                                   +---------------------+
                                                   | movie_ID (PK)       |
                                                   | title               |
                                                   | type                |
                                                   | runtime             |
                                                   | release_date        |
                                                   | director            |
                                                   | starring            |
                                                   +---------------------+
                                                          /|\
                                                           |
                                                   shown_in|
                                                           |
                                                          \|/
                                                   +---------------------+
                                                   |        hall         |
                                                   +---------------------+
                                                   | hall_ID (PK)        |
                                                   | mode                |
                                                   | capacity            |
                                                   | location            |
                                                   +---------------------+
