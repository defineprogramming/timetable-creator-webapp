import (
	"database/sql"
	"time"
	"web-app/utils"

	"github.com/golang-jwt/jwt" // Import JWT package (requires github.com/golang-jwt/jwt package)
)

// User represents a registered user 
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login validates user credentials and returns a JWT token 
func (u *User) Login() (string, error) {

	db, err := sql.Open("sqlite3", "./web-app.db") // Open database connection (requires sqlite3 driver package)
	if err != nil {
		return "", err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, name FROM users WHERE email = ? AND password = ?", u.Email, u.Password) // Query database for user by email and password

	err = row.Scan(&u.ID, &u.Name) // Scan row into user struct 
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256) // Create a new token with HS256 signing method

	claims := token.Claims.(jwt.MapClaims) // Cast token claims to map claims

	claims["id"] = u.ID                   // Set user ID as claim
	claims["name"] = u.Name               // Set user name as claim
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set expiration time as claim

	tkString, err := token.SignedString([]byte(utils.JWTSecret)) // Sign token with secret key (defined in utils/config.go)
	if err != nil {
		return "", err
	}

	return tkString,nil 

}

// GetUserByToken returns a user for a given JWT token from the database 
func GetUserByToken(tokenString string) (User,error){

	varuserUser 

	token,err:=jwt.Parse(tokenString ,func(token*jwt.Token)(interface{},error){
			return[]byte(utils.JWTSecret),nil// Parse token with secret key (defined in utils/config.go)
     })
	iferr!=nil{
			returnuser,err 
     }

	if!token.Valid{// Check if token is valid 
			returnuser,jwt.ErrInvalidToken 
     }

	mapClaims:=token.Claims.(jwt.MapClaims)// Cast token claims to map claims 

	userID:=int64(mapClaims["id"].(float64))// Get user ID from claim 

	db,err:=sql.Open("sqlite3","./web-app.db")// Open database connection (requires sqlite3 driver package)
	iferr!=nil{
			returnuser,err 
     }
	deferdb.Close()

	row:=db.QueryRow("SELECT id,name,email FROM users WHERE id = ?",userID)// Query database for user by ID 

	err=row.Scan(&user.ID,&user.Name,&user.Email)// Scan row into user struct 
	iferr!=nil{
			returnuser,err 
     }

	returnuser,nil 

}
