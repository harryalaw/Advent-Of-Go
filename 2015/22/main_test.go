package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		player Player
		boss   Boss
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4 HP Boss",
			args: args{
				player: Player{
					hp:   10,
					mana: 250,
				},
				boss: Boss{
					hp:     4,
					damage: 8,
				},
			},
			want: 53,
		},
		{
			name: "6 HP Boss, 8 damage",
			args: args{
				player: Player{
					hp:   10,
					mana: 250,
				},
				boss: Boss{
					hp:     6,
					damage: 8,
				},
			},
			want: 106,
		},
		{
			name: "6 HP Boss, 11 damage",
			args: args{
				player: Player{
					hp:   10,
					mana: 250,
				},
				boss: Boss{
					hp:     6,
					damage: 11,
				},
			},
			want: 126,
		},
		{
			name: "13 HP Boss, 8 damage",
			args: args{
				player: Player{
					hp:   10,
					mana: 250,
				},
				boss: Boss{
					hp:     13,
					damage: 8,
				},
			},
			want: 226,
		},
		{
			name: "14 HP Boss, 8 damage",
			args: args{
				player: Player{
					hp:   10,
					mana: 250,
				},
				boss: Boss{
					hp:     14,
					damage: 8,
				},
			},
			want: 641,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.player, tt.args.boss); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		player Player
		boss   Boss
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5 HP Boss, 9 damage",
			args: args{
				player: Player{
					hp:   10,
					mana: 250,
				},
				boss: Boss{
					hp:     5,
					damage: 9,
				},
			},
			want: 126,
		},
		{
			name: "8 HP Boss, 8 damage",
			args: args{
				player: Player{
					hp:   10,
					mana: 250,
				},
				boss: Boss{
					hp:     8,
					damage: 8,
				},
			},
			want: 219,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.player, tt.args.boss); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
