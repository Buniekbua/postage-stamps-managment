
<div id="top"></div>


[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />


<h3 align="center">Postage Stamps Managment</h3>

  <p align="center">
    <a href="https://github.com/Buniekbua/postage-stamps-managment"><strong>Explore the docs »</strong></a>
    <br />
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
     <li>
      <a href="#structure">Structure</a>
      <ul>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project
Project is a RESTful API for people who want store their postage stamps album in the app. 
App is written in Go using **Go Fiber** framework, **GORM** library and **JWT** authentication.

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [![Go][Go.dev]][Go-url]
*  [![Fiber][Gofiber.io]][Fiber-url]
* [![	Gorm][Gorm.io]][Gorm-url]
* [![Sqlite][Sqlite]][Sqlite-url]
* [![Jwt][Jwt]][Jwt-url]

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started


### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/Buniekbua/postage-stamps-managment.git
   ```
2. Add .env file with **SECRET_KEY**
```sh
SECRET_KEY = YourSecretKey
```
4. Build
   ```sh
   cd postage-stamps-managment
   go build
   ```
5. Run
   ```sh
   ./postage-stamps-managment
   ```

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- STRUCTURE -->
## Structure

  
  📦postage-stamps-managment
 ┣ 📂controllers  
 ┃ ┣ 📜authController.go  
 ┃ ┣ 📜stampController.go  
 ┃ ┗ 📜userController.go  
 ┣ 📂database  
 ┃ ┗ 📜database.go  
 ┣ 📂middleware  
 ┃ ┗ 📜authMiddleware.go  
 ┣ 📂models  
 ┃ ┣ 📜stampModel.go  
 ┃ ┗ 📜userModel.go  
 ┣ 📂routes  
 ┃ ┣ 📜authRoutes.go  
 ┃ ┣ 📜stampRoutes.go  
 ┃ ┗ 📜userRoutes.go  
 ┗ 📜main.go
  
<!-- API -->
## API
**/users**
* ``GET``: Get all users (login needed)

**/users:id**
* ``GET``: Get a user (login needed)
* ``DELETE``: Delete a user (login needed)

**/stamps**
* ``GET``: Get all stamps
* ``POST``: Create a new stamp (login needed)

**/stamps/:id**
* ``GET``: Get a stamp
* ``DELETE``: Delete a stamp (login needed)
* ``PATCH``: Update a stamp (login needed)

**/signup**
* ``POST``: Create a new user

**/login**
* ``POST``: Login a user

**/validate**
* ``GET``: Validate a user (login needed)





<!-- CONTACT -->
## Contact

Buniekbua  - buniewicz.szymon@gmail.com

Project Link: [https://github.com/Buniekbua/postage-stamps-managment](https://github.com/Buniekbua/postage-stamps-managment)

<p align="right">(<a href="#top">back to top</a>)</p>







<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/szymon-buniewicz-612459238
[product-screenshot]: images/screenshot.png
[Go.dev]:https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]:https://go.dev/
[Gofiber.io]:https://img.shields.io/badge/Fiber-08add5?style=for-the-badge&logo=fiber&logoColor=white
[Fiber-url]:https://gofiber.io/
[Gorm.io]: https://img.shields.io/badge/Gorm-157ad3?style=for-the-badge&logo=gorm&logoColor=white
[Gorm-url]: https://gorm.io/
[Sqlite]: https://img.shields.io/badge/Sqlite-003B57?style=for-the-badge&logo=sqlite&logoColor=white
[Sqlite-url]: https://www.sqlite.org/
[Jwt]: https://img.shields.io/badge/Jwt-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white
[Jwt-url]: https://jwt.io/
