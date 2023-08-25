# server-core
The main golang server of the game

## PLans
### v0.1.0
Server must be able to accept json, what can request some data from server, also say like "000.000.0.0:6000 - moved right", and send requests with world map on start, send data about that changes, and (example) if player sayed move to the right, it must send data about allowing or not allowing this. also it need to store information about active and connected in this session users, and must be able to store long-time information about player in json. Data stored in different files in file system of server, but client only says like "give me the font", and server store information what font is locaated on this path. Also some info like graphics must be decoded from (for example) .png to an array of colors (colors stored in format of u32 (unsigned integer of 32 bit) like 0xAARRGGBB), and in request must be sended that array, embedded into field in response json. Also server must be able to send data about players movement to all other players.
