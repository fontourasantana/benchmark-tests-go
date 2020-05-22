package repositories

import (
	"fmt"
	"encoding/json"
	"ameicosmeticos/app/contracts"
	"ameicosmeticos/app/models"
	_ "github.com/go-sql-driver/mysql"
)

type CanceledUsersRepository struct {
	db		contracts.IDbHandler
	cache	contracts.ICacheHandler
}

const (
	redisExpire = 60
)

func NewCanceledUsersRepository(db contracts.IDbHandler, cache contracts.ICacheHandler) *CanceledUsersRepository {
	return &CanceledUsersRepository{
		db,
		cache,
	}
}

func (this *CanceledUsersRepository) GetUser(id uint64) (models.CanceledUsers, error) {
	var user models.CanceledUsers
	key := fmt.Sprintf("getuser_%d", id);
	data, err := this.cache.Get(key)

	if err != nil {
		row, err := this.db.Query(fmt.Sprintf("SELECT * FROM canceled_users WHERE id = %d", id))

		if err != nil {
			return models.CanceledUsers{}, err
		}

		row.Next()
		row.Scan(&user.ID, &user.UserId, &user.Username, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		row.Close()

		serializedUser, err := json.Marshal(user)
		if err != nil {
			fmt.Println("erro ao serializar objeto: ", err)
		}

		this.cache.Set(key, serializedUser) // Adicionar um erro interno para identificar que o cache está com problema

		return user, nil
	}
	
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println(err)
	}

	return user, nil
}

// type PlayerRepository struct {
// 	interfaces.IDbHandler
// }

// func (repository *PlayerRepository) GetPlayerByName(name string) (models.PlayerModel, error) {

// 	row, err :=repository.Query(fmt.Sprintf("SELECT * FROM player_models WHERE name = '%s'", name))
// 	if err != nil {
// 		return models.PlayerModel{}, err
// 	}

// 	var player models.PlayerModel

// 	row.Next()
// 	row.Scan(&player.Id, &player.Name, &player.Score)

// 	return player, nil
// }