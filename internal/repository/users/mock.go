package users

import (
	"github.com/stretchr/testify/mock"

	"github.com/Kiritogtsa/pi_in_go/internal/domains/entries"
)

// Estrutura que implementa a interface Userif com o mock de testify
type mockuser struct {
	mock.Mock
	data   map[int]*entries.User // Simulação do banco de dados
	nextID int                   // Controle do próximo ID
}

// Função para criar um novo mockuser
func Newmock() Userif {
	return &mockuser{
		data:   make(map[int]*entries.User), // Inicializa o mapa simulado
		nextID: 1,                           // Inicializa o próximo ID
	}
}

// Método de atualização de um usuário no mock
func (m *mockuser) update(user *entries.User) error {
	if _, exists := m.data[user.ID]; !exists {
		return &NotFoundError{ID: user.ID}
	}
	m.data[user.ID] = user // Atualiza o usuário na simulação do banco
	args := m.Called(user)
	return args.Error(0)
}

// Método para deletar um usuário pelo ID
func (m *mockuser) Delete(id int) error {
	if _, exists := m.data[id]; !exists {
		return &NotFoundError{ID: id}
	}
	delete(m.data, id) // Remove o usuário da simulação do banco
	args := m.Called(id)
	return args.Error(0)
}

// Método para retornar todos os usuários
func (m *mockuser) Getall() ([]entries.User, error) {
	var users []entries.User
	for _, user := range m.data {
		users = append(users, *user) // Constrói a lista de usuários da simulação
	}
	args := m.Called()
	return users, args.Error(0)
}

// Método para persistir um usuário
func (m *mockuser) Persit(user *entries.User) error {
	if user.ID == 0 {
		return m.insert(user)
	} else {
		return m.update(user)
	}
}

// Método para inserir um novo usuário
func (m *mockuser) insert(user *entries.User) error {
	user.ID = m.nextID     // Gera um novo ID
	m.nextID++             // Incrementa o próximo ID
	m.data[user.ID] = user // Insere o usuário na simulação do banco
	args := m.Called(user)
	return args.Error(0)
}

// NotFoundError define um erro personalizado para quando um usuário não é encontrado
type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return "user not found"
}
func (m *mockuser) GetByemail(string) (*entries.User, error) {
	return nil, nil
}
