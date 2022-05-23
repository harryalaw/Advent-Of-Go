package main

import "testing"

func Test_playerWinsFight(t *testing.T) {
	type args struct {
		player Player
		boss   Player
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Player wins",
			args: args{
				player: Player{hp: 8, damage: 5, armor: 5},
				boss:   Player{hp: 12, damage: 7, armor: 2},
			},
			want: true,
		},
		{name: "Boss wins",
			args: args{
				player: Player{hp: 12, damage: 7, armor: 2},
				boss:   Player{hp: 8, damage: 5, armor: 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playerWinsFight(tt.args.player, tt.args.boss); got != tt.want {
				t.Errorf("playerWinsFight() = %v, want %v", got, tt.want)
			}
		})
	}
}
