package models

import (
	"time"
	"gorm.io/gorm"
)

type Usuario struct {
	ID           uint           `gorm:"primaryKey;column:id" json:"id"`
	Nome         string         `gorm:"column:nome" json:"nome"`
	Email        string         `gorm:"column:email;uniqueIndex" json:"email"`
	Senha        string         `gorm:"column:senha" json:"-"`
	CriadoEm     time.Time      `gorm:"column:criado_em" json:"criado_em"`
	AtualizadoEm time.Time      `gorm:"column:atualizado_em" json:"atualizado_em"`
	DeletadoEm   gorm.DeletedAt `gorm:"column:deletado_em;index" json:"-"`
}

func (Usuario) tableName() string {
	return "usuario"
}