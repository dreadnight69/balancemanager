package config

const (
	dbHostEnv     = "DB_HOST"
	dbPortEnv     = "DB_PORT"
	dbNameEnv     = "DB_NAME"
	dbUserEnv     = "DB_USER"
	dbPasswordEnv = "DB_PASSWORD"
)

type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (p *DB) Init() (err error) {
	p.Host, err = getStringFromEnv(dbHostEnv)
	if err != nil {
		return err
	}

	p.Port, err = getStringFromEnv(dbPortEnv)
	if err != nil {
		return err
	}

	p.Name, err = getStringFromEnv(dbNameEnv)
	if err != nil {
		return err
	}

	p.User, err = getStringFromEnv(dbUserEnv)
	if err != nil {
		return err
	}

	p.Password, err = getStringFromEnv(dbPasswordEnv)

	return err
}
