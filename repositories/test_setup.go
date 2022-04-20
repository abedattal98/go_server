package repositories
import(
	"rgb/domain"
	"rgb/models"
	"time"
)
func AddTestUser(repo domain.IUserRepository) (models.User, error) {
	user := models.User{
		Username: "rand.String(10)",
		Password: "secret123",
		ID:       int(time.Now().Unix()),
	}
	user, err := repo.Save(user)
	return user, err
}