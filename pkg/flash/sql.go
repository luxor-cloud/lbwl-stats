package flash


type SQLPlayerRepository struct {

}

func (spr *SQLPlayerRepository) GetByUUID(uuid string) (*Player, error) {

}

func (spr *SQLPlayerRepository) Update(player *Player) error {

}

func (spr *SQLPlayerRepository) Exists(uuid string) (bool, error) {

}
