package myfunctions


import (
	"errors"
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)


func NewPlayer(name string)(p mt.Player){
	p.Name = name
	return p
}

func NewTeam(name string)(t mt.Team) {
	t.Name = name
	return t
}

func isPlayerInTeam(p mt.Player, t mt.Team) (bool){
	if p.Teams == nil {
		return false
	}
	for i := range t.PlayerList {
		if t.PlayerList[i].Name == p.Name {
			return true
		}
	}
	return false
}

func isTeamInPlayerTeams(p mt.Player, t mt.Team) (bool){
	if p.Teams == nil {
		return false
	}
	for i := range p.Teams {
		if p.Teams[i].Name == t.Name {
			return true
		}
	}
	return false
}

func AddPlayerToTeam(p *mt.Player, t *mt.Team)(err error) {
	if isPlayerInTeam(*p, *t) {
		return errors.New(fmt.Sprintf("%v is already in team %v", p.Name, t.Name))
	}
	p.Teams = append(p.Teams, t)
	t.PlayerList = append(t.PlayerList, p)
	return nil
}

func RemovePlayerFromTeam(p *mt.Player, t *mt.Team)(err error) {
	if !isPlayerInTeam(*p, *t) {
		return errors.New(fmt.Sprintf("%v is not in team %v", p.Name, t.Name))
	} 
	for i := range t.PlayerList {
		if t.PlayerList[i].Name == p.Name {
			t.PlayerList[i] = t.PlayerList[len(t.PlayerList)-1]
			t.PlayerList = t.PlayerList[:len(t.PlayerList)-1]
		}
	}
	for i := range p.Teams {
		if p.Teams[i].Name == t.Name {
			p.Teams[i] = p.Teams[len(p.Teams)-1]
			p.Teams = p.Teams[:len(p.Teams)-1]
		}
	}
	return nil
}

func DeletePlayer(p *mt.Player)(err error) {
	pName := p.Name
	if p.Teams != nil {
		for _, team := range p.Teams {
			RemovePlayerFromTeam(p, &team)
		}
	}
	p = nil
	return errors.New(fmt.Sprintf("%f has been successfully deleted", pName))
}

func DeleteTeam(t *mt.Team)(err error) {
	tName := t.Name
	if t.PlayerList != nil {
		for _, player := range t.PlayerList {
			RemovePlayerFromTeam(&player, t)
		}
	}
	t = nil
	return errors.New(fmt.Sprintf("%f has been successfully deleted", tName))
}
