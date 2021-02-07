package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var STDIN io.Reader

// Declarations in BBC BASIC are global so must be outside the functions

// 30
var D, N, C, A, O int
var C_, X_, B_ string

// 670
var A_, D_ string

// 2820
var LR, Z, V, W, C4 int
var I_ = new([8]string)[0:8]
var L = new([53]int)[0:53]
var F = new([53]int)[0:53]

// 4130
// BBC BASIC allows two variables to have the same name if one is a simple variable and one an array variable
// Modern languages do not allow this, so I've converted L->LL and F->FF
var R /*B,*/, LL, E, C1, C2, C3, FF, G, Y, X int
var H_, Q_, G_, F_, K_, L_, M_, N_, U_, W_, P_, R_, S_, T_, J_, O_, V_, Z_ string

func main() {

	STDIN = os.Stdin
	SEED := time.Now().UnixNano()

	var initialCommands []byte
	if len(os.Args) > 1 {
		commandFile := os.Args[1]
		cfh, err := os.Open(commandFile)
		if err != nil {
			fmt.Printf("Error opening file: %s", err.Error())
			os.Exit(1)
		}
		defer cfh.Close()
		initialCommands, err = ioutil.ReadAll(cfh)
		if err != nil {
			fmt.Printf("Error reading file: %s", err.Error())
			os.Exit(1)
		}
		// close now so we don't leave it open for a potentially long game run
		cfh.Close()
	}
	if len(os.Args) > 2 {
		seedArg := os.Args[2]
		seed, seedErr := strconv.Atoi(seedArg)
		if seedErr != nil {
			fmt.Printf("Error parsing random number seed: %s", seedErr.Error())
			os.Exit(2)
		}
		SEED = int64(seed)
	}

	if len(initialCommands) > 0 {
		STDIN = io.MultiReader(bytes.NewReader(initialCommands), STDIN)
	}

	rand.Seed(SEED)

	// 10REM ISLAND OF SECRETS
	GOSUB_2820()

label_30:
	D = R
	if R == 20 {
		D = FNR(80)
	}

	GOSUB_650()
	GOSUB_2770()
	PRINT("ISLAND OF SECRETS   TIME REMAINING:" + STR_(LL))
	PRINT_(G_)
	TAB(0)
	PRINT_("STRENGTH = " + STR_(Y))
	TAB(23)
	PRINT("WISDOM = " + STR_(X))
	PRINT(G_)
	PRINT_("YOU ARE " + I_[VAL(LEFT_(A_, 1))] + " ")
	GOSUB_720()
	N = 0
	for I := 1; I <= C4; I++ {
		C = 0
		Y_ := READ()
		if L[I] == R && F[I] < 1 {
			C = 1
		}
		if N > 0 && C == 1 {
			A_ = A_ + ","
		}
		if C == 1 {
			A_ = A_ + " " + Y_
			N = N + 1
		}
	}
	if N > 0 {
		A_ = "*YOU SEE" + A_
		GOSUB_720()
	}
	PRINT("")
	PRINT(G_)
	PRINT(F_)
	PRINT("")
	PRINT_("WHAT WILL YOU DO")
	E_ := INPUT(STDIN)
	C_ = ""
	X_ = ""
	A = 0
	O = 52
	LI := LEN(E_)
	for I := 1; I <= LI; I++ {
		if MID_(E_, I, 1) == " " && C_ == "" {
			C_ = LEFT_(E_, I-1)
		}
		if MID_(E_, I+1, 1) != " " && C_ > "" {
			X_ = RIGHT_(E_, LI-I)
			I = LI
		}
	}
	if X_ == "" {
		C_ = E_
	}
	if LEN(C_) < 3 {
		C_ = C_ + "???"
	}
	for I := 1; I <= V; I++ {
		if LEFT_(C_, 3) == MID_(V_, 3*(I-1)+1, 3) {
			A = I
		}
	}
	GOSUB_760()

	B_ = ""
	if A == 0 {
		A = V + 1
	}
	if X_ == "???" {
		F_ = "MOST ACTIONS NEED TWO WORDS"
	}
	if A > V || O == 52 {
		F_ = W_ + C_ + " " + X_
	}
	if A > V && O == 52 {
		F_ = "WHAT!"
	}
	LL = LL - 1
	Y = FNS(Z)
	B_ = STR_(O) + STR_(L[O]) + STR_(F[O]) + STR_(R)
	ON_GOSUB(INT(float64(A)/10)+1, GOSUB_590, GOSUB_600, GOSUB_610, GOSUB_620, GOSUB_630, GOSUB_630, GOSUB_640)
	if R == 61 {
		X = X - FNR(2) + 1
	}
	if R == 14 && FNR(3) == 1 {
		Y = Y - 1
		F_ = "YOU ARE BITTEN"
	}
	if F[36] < 1 && -R != F[22] {
		F[36] = F[36] + 1
		L[36] = R
		Y = Y - 1
	}
	if R != L[16] && L[16] > 0 {
		L[16] = 1 + FNR(4)
	}
	if R != L[39] {
		L[39] = 10*(FNR(5)-1) + 7 + FNR(3)
	}
	if R == L[39] && R != L[43] && F[13] > -1 {
		Y = Y - 2
		X = X - 2
	}
	if R < 78 {
		L[32] = 76 + FNR(2)
	}
	if R == 33 || R == 57 || R == 73 && FNR(2) == 1 {
		L[25] = R
	}
	if R == L[32] && FNR(2) == 1 && F[32] == 0 {
		GOSUB_1310()
	}
	if R == 19 && Y < 70 && F[43] == 0 && FNR(4) == 1 {
		F_ = "PUSHED INTO THE PIT"
		F[W] = 1
	}
	if R != L[41] {
		L[41] = 21 + (FNR(3) * 10) + FNR(2)
	}
	if R == L[41] {
		F[41] = F[41] - 1
		if F[41] < -4 {
			GOSUB_1230()
		}
	}
	if F[43] == 0 {
		L[43] = R
	}
	if L[43] < 18 && R != 9 && R != 10 && F[W-2] < 1 {
		GOSUB_1330()
	}
	if R == 18 {
		Y = Y - 1
	}
	if Y < 50 {
		O = FNR(9)
		GOSUB_1530()
		if L[O] == R {
			F_ = "YOU DROP SOMETHING"
		}
	}
	if LL < 900 && R == 23 && F[36] > O && FNR(3) == 3 {
		GOSUB_1360()
	}
	if R == 47 && F[8] > 0 {
		F_ = F_ + " YOU CAN GO NO FURTHER"
	}
	if F[8]+F[11]+F[13] == -3 {
		F[W] = 1
		GOSUB_2800()
	}
	if F[W] == 0 && LL > 0 && Y > 1 && X > 1 {
		goto label_30
	}
	if LL < 1 || Y < 1 {
		F_ = "YOU HAVE FAILED, THE EVIL ONE SUCCEEDS"
	}
	PRINT("")
	PRINT(F_)
	PRINT("YOUR FINAL SCORE=" + STR_(X+Y+(ABS(LL/7*(BOOL2INT(LL < 640))))))
	PRINT("")
	PRINT("")
	PRINT("GAME OVER")
	END()
}

func GOSUB_590() {
	ON_GOSUB(A, GOSUB_810, GOSUB_810, GOSUB_810, GOSUB_810, GOSUB_810, GOSUB_1080, GOSUB_1080, GOSUB_1390, GOSUB_1530)
}

func GOSUB_600() {
	ON_GOSUB(A-9, GOSUB_1540, GOSUB_1630, GOSUB_1670, GOSUB_1710, GOSUB_1730, GOSUB_1080, GOSUB_1760, GOSUB_1760, GOSUB_1760, GOSUB_1760)
}

func GOSUB_610() {
	ON_GOSUB(A-19, GOSUB_1820, GOSUB_1820, GOSUB_1820, GOSUB_1820, GOSUB_1910, GOSUB_2100, GOSUB_2210, GOSUB_2270, GOSUB_2270, GOSUB_1080)
}

func GOSUB_620() {
	ON_GOSUB(A-29, GOSUB_2500, GOSUB_2500, GOSUB_2300, GOSUB_2300, GOSUB_2330, GOSUB_2350, GOSUB_2400, GOSUB_2400, GOSUB_2470, GOSUB_2540)
}
func GOSUB_630() {
	ON_GOSUB(A-39, GOSUB_2600, GOSUB_2600, GOSUB_2720, GOSUB_640)
}

func GOSUB_640() {
}

func GOSUB_650() {
	D = D*10 + LR
	DATA_POINTER = (D - 2870) / 10
	A_ = READ()
	D_ = RIGHT_(A_, 4)
	A_ = LEFT_(A_, LEN(A_)-4)
	if R == 39 {
		D_ = MID_("101110100", FNR(5), 4)
	}
	if R == 20 {
		D_ = "1110"
	}
	GOSUB_2780()
}

func GOSUB_720() {
	for I := 2; I <= LEN(A_); I++ {
		E_ := MID_(A_, I, 1)
		PRINT_(E_)
		if E_ == " " && FNP(Z) > Z {
			PRINT("")
		}
	}
	PRINT_(". ")
	A_ = ""
}

func GOSUB_760() {
	if LEN(X_) < 3 {
		X_ = X_ + "???"
	}
	for I := 1; I <= W; I++ {
		if LEFT_(X_, 3) == MID_(Z_, 3*(I-1)+1, 3) {
			O = I
		}
	}
	if O == 0 {
		O = 52
	}
}

func GOSUB_810() {
	D = 0
	C = 0
	if O == 52 {
		D = A
	}
	if O > C4 && O < W {
		D = O - C4
	}
	if B_ == "500012" || B_ == "500053" || B_ == "500045" {
		D = 4
	}
	if B_ == "500070" || B_ == "500037" || B_ == "510011" || B_ == "510041" {
		D = 1
	}
	if B_ == "510043" || B_ == "490066" || B_ == "490051" {
		D = 1
	}
	if B_ == "510060" || B_ == "480056" {
		D = 2
	}
	if B_ == "510044" || B_ == "510052" {
		D = 3
	}
	if B_ == "490051" && F[29] == 0 {
		GOSUB_2110()
		return
	}
	if R == L[39] && (X+Y < 180 || R == 10) {
		F_ = W_ + "LEAVE!"
		return
	}
	if R == L[32] && F[32] < 1 && D == 3 {
		F_ = "HE WILL NOT LET YOU PAST"
		return
	}
	if R == 47 && F[44] == 0 {
		F_ = "THE ROCKS MOVE TO PREVENT YOU"
		return
	}
	if R == 28 && F[7] != 1 {
		F_ = "THE ARMS HOLD YOU FAST"
		return
	}
	if R == 45 && F[40] == 0 && D == 4 {
		F_ = "HISSSS!"
		return
	}
	if R == 25 && F[16]+L[16] != -1 && D == 3 {
		F_ = "TOO STEEP TO CLIMB"
		return
	}
	if R == 51 && D == 3 {
		F_ = "THE DOOR IS BARRED!"
		return
	}
	if D > 0 && MID_(D_, D, 1) == "0" {
		R = R + VAL(MID_("-10+10+01-01", D*3-2, 3))
		C = 1
	}
	F_ = "OK"
	if D < 1 || C == 0 {
		F_ = W_ + "GO THAT WAY"
	}
	if R == 33 && L[16] == 0 {
		L[16] = FNR(4)
		F[16] = 0
		F_ = "THE BEAST RUNS AWAY!"
	}
	if R != L[25] || O != 25 {
		return
	}
	F_ = ""
	A_ = "#YOU BOARD THE CRAFT "
	if X < 60 {
		A_ = A_ + S_
	}
	A_ = A_ + T_
	GOSUB_2740()
	GOSUB_2760()
	GOSUB_2760()
	if X < 60 {
		A_ = "#TO SERVE ONEGAN FOREVER!"
		F[W] = 1
	}
	if X > 59 {
		A_ = "#THE BOAT SKIMS THE DARK SILENT WATERS"
		R = 57
	}
	GOSUB_2750()
	GOSUB_2760()
	GOSUB_2760()
}

func GOSUB_1080() {
	if ((F[O] > 0 && F[O] < 9) || L[O] != R) && O <= C3 {
		F_ = "WHAT " + X_ + "?"
		return
	}
	if B_ == "3450050" {
		Y = Y - 8
		X = X - 5
		F_ = "THEY ARE CURSED"
		return
	}
	if B_ == "3810010" {
		GOSUB_1370()
	}
	if (A == 15 && O != 20 && O != 1) || (A == 29 && O != 16) || O > C3 {
		F_ = W_ + C_ + " " + X_
		return
	}
	if L[O] == R && (F[O] < 1 || F[O] == 9) && O < C3 {
		L[O] = 0
		A = -1
	}
	if O == 16 && L[10] != 0 {
		L[O] = R
		F_ = "IT ESCAPED"
		A = 0
	}
	if O > C1 && O < C2 {
		FF = FF + 2
		A = -1
	}
	if O >= C2 && O <= C3 {
		G = G + 2
		A = -1
	}
	if O > C1 && O < C3 {
		L[O] = -81
	}
	if A == -1 {
		F_ = "TAKEN"
		X = X + 4
		E = E + 1
		if F[O] > 1 {
			F[O] = 0
		}
	}
	if B_ != "246046" || L[11] == 0 {
		return
	}
	F_ = U_
	L[O] = R
	if FNR(3) < 3 {
		return
	}
	GOSUB_1200()
}

func GOSUB_1200() {
	A_ = "#" + U_ + R_
	R = 63 + FNR(6)
	L[16] = 1
	F_ = ""
	GOSUB_2740()
}

func GOSUB_1230() {
	GOSUB_2770()
	F_ = ""
	A_ = "#THE LOGMEN " + M_
	F[41] = 0
	Y = Y - 4
	X = X - 4
	if R < 34 {
		A_ = A_ + "THROW YOU IN THE WATER"
		R = 32
	}
	if R > 33 {
		A_ = A_ + "TIE YOU UP IN A STOREROON"
		R = 51
	}
	GOSUB_2750()
	GOSUB_2760()
	for I := 3; I <= 4; I++ {
		if L[I] == 0 {
			L[I] = 42
		}
	}
}

func GOSUB_1310() {
	A_ = "*THE SWAMPMAN TELLS HIS TALE"
	GOSUB_2740()
	F[32] = -1
}

func GOSUB_1330() {
	F_ = "MEDIAN CAN DISABLE THE EQUIPMENT"
	if L[8] == 0 {
		F_ = F_ + " AND ASKS YOU FOR THE PEBBLE YOU CARRY"
	}
}

func GOSUB_1360() {
	F[36] = -(FNR(4) + 6)
	F_ = "A STORM BREAKS OVERHEAD!"
}

func GOSUB_1370() {
	for K := 1; K <= 30; K++ {
		GOSUB_2770()
		PRINT("////LIGHTNING FLASHES!")
	}
	L[39] = R
	Y = Y - 8
	X = X - 2
}

func GOSUB_1390() {
	if (O != 24 && L[O] > 0) || O == 52 {
		F_ = "YOU DON'T HAVE THE " + X_
		return
	}

	PRINT_("GIVE THE " + X_ + " TO WHOM")
	X_ = INPUT(STDIN)
	Q := O
	GOSUB_760()
	N = O
	O = Q
	if R != L[N] {
		F_ = "THE " + X_ + " IS NOT HERE"
		return
	}
	if B_ == "10045" && N == 40 {
		L[O] = 81
		F[40] = 1
		F_ = "THE SNAKE UNCURLS"
	}
	if B_ == "2413075" && N == 30 && G > 1 {
		F[11] = 0
		F_ = "HE OFFERS HIS STAFF"
		G = G - 1
	}
	B_ = LEFT_(B_, 3)
	F_ = "IT IS REFUSED"
	if B_ == "300" && N == 42 {
		X = X + 10
		L[O] = 81
	}
	if B_ == "120" && N == 42 {
		X = X + 10
		L[0] = 81
	}
	if B_ == "40-" && N == 32 {
		F[N] = 1
		L[O] = 81
	}
	if LEFT_(B_, 2) == "80" && N == 43 {
		L[O] = 81
		GOSUB_1560()
	}
	if L[O] == 81 || (O == 24 && L[11] > 0 && G > 0) {
		F_ = "IT IS ACCEPTED"
	}
	if N == 41 {
		L[O] = 51
		F_ = "IT IS TAKEN"
	}
}

func GOSUB_1530() {
	if O == 4 && L[O] == 0 {
		L[O] = 81
		X = X - 1
		F_ = "IT BREAKS!"
		return
	}
	GOSUB_1540()
}

func GOSUB_1540() {
	if L[O] == 0 && O <= C1 {
		L[O] = R
		F_ = "DONE"
		E = E - 1
	}
}

func GOSUB_1560() {
	A_ = "*HE TAKES IT "
	if R != 8 {
		A_ = A_ + "RUNS DONN THE CORRIDOR, "
	}
	GOSUB_2740()
	A_ = "*AND CASTS IT INTO THE CHEMICAL VATS, PURIFYING THEM WITH"
	A_ = A_ + " A CLEAR BLUE LIGHT REACHING FAR INTO THE LAKES AND RIVERS BEYOND"
	F[8] = -1
	GOSUB_2750()
	GOSUB_2760()
	GOSUB_2760()
}

// This appears to be uncalled code in the original source!
// 16001F L[I]!=0 && I<C1 THEN LET I=1+1: GOTO1600
// 16101F L[I]==0 THEN LET L[I]=R:LET F[I]=0: GOSUB1540:LET F_="YOU DROP SOMETHING"
// 1620RETURN

func GOSUB_1630() {
	if (O < C1 || O > C3) && X_ != "???" {
		F_ = W_ + C_ + " " + X_
		X = X - 1
		return
	}
	F_ = "YOU HAVE NO FOOD"
	if FF > 0 {
		FF = FF - 1
		Y = Y + 10
		F_ = "OK"
	}
	if O == 3 {
		X = X - 5
		Y = Y - 2
		F_ = "THEY MAKE YOU VERY ILL!"
	}
}

func GOSUB_1670() {
	if O == 31 {
		GOSUB_2380()
		return
	}
	if X_ != "???" && (O < 21 || O > C3) {
		F_ = W_ + C_ + " " + X_
		X = X - 1
		return
	}
	F_ = "YOU HAVE NO DRINK"
	if G > 0 {
		G = G - 1
		Y = Y + 7
		F_ = "OK"
	}
}

func GOSUB_1710() {
	if LEFT_(B_, 4) == "1600" {
		F[O] = -1
		F_ = "IT ALLOWS YOU TO RIDE"
	}
}
func GOSUB_1730() {
	if B_ == "2644044" {
		F_ = "CHEST OPEN"
		F[6] = 9
		F[5] = 9
		F[15] = 9
	}

	if B_ == "2951151" {
		F_ = "THE TRAPDOOR CREAKS"
		F[29] = 0
		X = X + 3
	}
}

func GOSUB_1760() {
	Y = Y - 2
	if B_ == "3577077" && L[9] == 0 {
		F[23] = 0
		L[23] = R
	}
	if V > 15 && V < 19 && (L[9] == 0 || L[15] == 0) {
		F_ = "OK"
	}
	if B_ == "1258158" || B_ == "2758158" && L[15] == 0 {
		F[12] = 0
		F[27] = 0
		F_ = "CRACK!"
	}
	if LEFT_(B_, 4) == "1100" && R == 10 {
		GOSUB_1980()
	}
	if A == 18 && (O > 29 && O < 34) || (O > 38 && O < 44) || O == 16 {
		GOSUB_1900()
	}
}

func GOSUB_1820() {
	Y = Y - 2
	X = X - 2
	if R != L[O] && L[O] != 0 {
		return
	}
	if O == 39 {
		F_ = "HE LAUGHS DANGEROUSLY"
	}
	if O == 32 {
		F_ = "THE SWAMPMAN IS UNMOVED"
	}
	if O == 33 {
		F_ = W_ + "TOUCH HER!"
		L[3] = 81
	}
	if O == 41 {
		F_ = "THEY THINK THAT'S FUNNY!"
	}
	if R == 46 {
		GOSUB_1200()
	}
	if LEFT_(B_, 4) == "1400" && R == L[39] {
		GOSUB_1980()
	}
	Y = Y - 8
	X = X - 5
}

func GOSUB_1900() {
	if L[9] > 0 {
		return
	}
	GOSUB_1910()
}

func GOSUB_1910() {
	Y = Y - 12
	X = X - 10
	F_ = "THAT WOULD BE UNWISE!"
	if R != L[O] {
		return
	}
	F[W] = 1
	A_ = "#THUNDER SPLITS THE SKY!"
	F_ = ""
	A_ = A_ + "#IT IS THE TRIUMPHANT VOICE OF OMEGAN."
	GOSUB_2740()
	A_ = "#WELL DONE ALPHAN! THE MEANS BECOME THE END.."
	A_ = A_ + "I CLAIM YOU AS MY OWN! HA HA HAH!"
	GOSUB_2750()
	GOSUB_2760()
	X = 0
	LL = 0
	Y = 0
}

func GOSUB_1980() {
	GOSUB_2770()
	ON_GOSUB(O-10, GOSUB_2010, GOSUB_2060, GOSUB_2060, GOSUB_2060)
	X = X + 10
	L[O] = 81
	F[O] = -1
	GOSUB_720()
	GOSUB_2760()
	GOSUB_2760()
}

func GOSUB_2010() {
	A_ = "#IT SHATTERS RELEASING A DAZZLING RAINBOW OF COLOURS!"
	if L[2] != R {
		return
	}
	A_ = A_ + "THE EGG HATCHES INTO A BABY DAKTYL " + O_
	L[39] = 81
	L[2] = 81
	F[2] = -1
	Y = Y + 40
}

func GOSUB_2060() {
	if L[13] != R {
		return
	}
	A_ = "*THE COAL BURNS WITH A WARM RED FLAME"
	F[13] = -1
	if R == 10 && R == L[39] {
		A_ = A_ + " WHICH DISSOLVES OMEGAN'S CLOAK"
		Y = Y + 20
	}
}

func GOSUB_2100() {
	if R != 51 || F[29] > 0 {
		F_ = W_ + C_ + " HERE"
		X = X + 1
	}
	GOSUB_2110()
}

func GOSUB_2110() {
label_2110:
	X = X - 1
	R = FNR(5)
	GOSUB_2770()
	PRINT("SWIMMING IN THE POISONOUS WATERS")
	J := 0
	B_ = ""
	F_ = "YOU SURFACE"
	PRINT("YOUR STRENGTH = " + STR_(Y))
	for I := 1; I <= R; I++ {
		if Y < 15 {
			PRINT("YOU ARE VERY WEAK")
		}
		PRINT_("WHICH WAY")
		X_ = INPUT(STDIN)
		X_ = LEFT_(X_, 1)
		B_ = B_ + X_
	}
	for I := 1; I <= R; I++ {
		Y = FNS(Z) - 3
		if MID_(B_, I, 1) == "N" {
			J = J + 1
		}
		if R/2 > J && Y > I {
			goto label_2110
		}
		if X < 2 {
			F_ = "YOU GOT LOST AND DROWNED"
		}
		R = 30 + FNR(3)
	}
}

func GOSUB_2210() {
	if F[36] > -1 {
		return
	}
	GOSUB_2770()
	PRINT("YOU CAN RUN TO SHELTER IN:")
	PRINT("1) GRANDPA'S SHACK")

	PRINT("2) CAVE OF SNELM")
	PRINT("3) LOG CABIN")
	PRINT("CHOOSE FROM 1-3")
	A_ = INPUT(STDIN)
	if A_ > "0" && A_ < "4" {
		R = ASC(MID_("A >", VAL(A_), 1)) - 21
		F[22] = -R
	}
	PRINT("YOU RUN BLINDLY THROUGH THE STORM")
	F_ = "YOU REACH SHELTER"
	GOSUB_2760()
}
func GOSUB_2270() {
	if B_ == "3075075" || B_ == "3371071" {
		F_ = "HOW WILL YOU DO THAT"
	}
	if B_ == "3371071" && A == 28 {
		F[3] = 0
		F_ = "SHE NODS SLOWLY"
		X = X + 5
	}
}

func GOSUB_2300() {
	F_ = "EXAMINE THE BOOK FOR CLUES"
	if LEFT_(B_, 3) == "600" {
		F_ = L_
	}
}

func GOSUB_2330() {
	if B_ == "40041" {
		F[4] = -1
		F_ = "FILLED"
	}
}

func GOSUB_2350() {
	F_ = X_
	if X_ == H_ && R == 47 && F[8] == 0 {
		F[44] = 1
		F_ = J_
	}
	if X_ != P_ || R != L[42] || L[3] < 81 || L[12] < 81 {
		return
	}
	F_ = "HE EATS THE FLOWERS- AND CHANGES"
	F[42] = 1
	F[43] = 0
}

func GOSUB_2380() {
	if F[4]+L[4] != -1 {
		F_ = "YOU DON'T HAVE " + X_
		return
	}
	GOSUB_2770()
	PRINT("YOU TASTE A DROP AND..")
	GOSUB_2760()
	F_ = "*OUCH!"
	Y = Y - 4
	X = X - 7
	GOSUB_2400()
}

func GOSUB_2400() {
	GOSUB_2770()
	for I := 1; I <= ABS(F[36]+3); I++ {
		LL = LL - 1
		if Y < 100 || -R == F[22] {
			Y = Y + 1
		}
		PRINT("TIME PASSES")
		GOSUB_2760()
	}
	if LL > 100 || F[36] < 1 {
		X = X + 2
		F[36] = 1
	}
	if A == 37 || A == 36 {
		F_ = "OK"
	}
}

func GOSUB_2470() {
	if R == L[25] {
		F_ = "THE BOATMAN WAVES BACK"
	}
	if LEFT_(B_, 3) == "700" {
		F[7] = 1
		F_ = N_
		X = X + 8
	}
}

func GOSUB_2500() {
	F_ = "A-DUB-DUB"
	if LEFT_(B_, 4) != "2815" {
		return
	}
	if F[O] == 1 {
		F[O] = 0
		F_ = K_
		return
	}
	if L[5] == 0 {
		F[8] = 0
		GOSUB_1080()
		F_ = "THE STONE UTTERS " + H_
	}
}

func GOSUB_2540() {
	GOSUB_2770()
	PRINT(" INFO - ITEMS CARRIED")
	GOSUB_2780()

	PRINT_(G_)
	TAB(0)
	PRINT_(" FOOD=" + STR_(FF))
	TAB(23)
	PRINT("DRINK=" + STR_(G))
	PRINT(G_)
	F_ = "OK"
	for I := 1; I <= C4; I++ {
		Y_ := READ()
		if L[I] == 0 {
			PRINT(Y_)
		}
	}
	PRINT(G_)
	GOSUB_2730()
}

func GOSUB_2600() {
	C_ = "LOAD"
	if A == 41 {
		C_ = "SAVE"
	}
	PRINT("PREPARE TO " + C_)
	GOSUB_2730()
	var FL *os.File
	if A == 40 {
		FL = OPENIN("ISDATA.SAVE")
	}
	if A == 41 {
		FL = OPENOUT("ISDATA.SAVE")
	}
	if A == 41 {
		F[50] = R
		F[49] = Y
		F[48] = X
		F[47] = FF
		F[46] = G
		F[45] = LL
	}
	for I := 1; I <= W; I++ {
		if A == 40 {
			L[I] = INPUT_DISC_NUMERIC(FL)
			F[I] = INPUT_DISC_NUMERIC(FL)
		}
		if A == 41 {
			PRINT_DISC_NUMERIC(FL, L[I])
			PRINT_DISC_NUMERIC(FL, F[I])
		}
	}
	CLOSE(FL)
	if A == 40 {
		R = F[50]
		Y = F[49]
		X = F[48]
		FF = F[47]
		G = F[46]
		LL = F[45]
	}
	F_ = "OK"
}

func GOSUB_2720() {
	F[W] = -1
	F_ = "YOU RELINQUISH YOUR QUEST."
	LL = 1
	return
}

func GOSUB_2730() {
	PRINT_("PRESS RETURN")
	A_ = INPUT(STDIN)
}

func GOSUB_2740() {
	GOSUB_2770()
	GOSUB_2750()
}

func GOSUB_2750() {
	GOSUB_720()
	GOSUB_2760()
}

func GOSUB_2760() {
	// I suspect this was some kind of delay...running on jsbeeb it takes about 1.5 seconds
	//for D := 1; D <= 2000; D++ {
	//}
	time.Sleep(1500 * time.Millisecond)
}

func GOSUB_2770() {
	CLS()
}

func GOSUB_2780() {
	DATA_POINTER = 80
}

func GOSUB_2800() {
	A_ = "*THE WORLD LIVES WITH NEW HOPE!"
	GOSUB_2750()
	F_ = "YOUR QUEST IS OVER"
}

func GOSUB_2820() {
	PRINT("INITIALISING")
	LR, Z = 2860, 39
	Z = INT(float64(Z) * 0.8)
	V, W, C4 = 42, 51, 43

	// lots of data statements

	// 4090
	GOSUB_2780()

	DATA_POINTER = 123
	for I := 1; I <= 7; I++ {
		I_[I] = READ()
	}
	R /*B,*/, LL, E = 23 /*8,*/, 1000, 0
	C1, C2, C3 = 16, 21, 24
	FF, G = 2, 2
	Y, X = 100, 35
	H_ = `MNgIL5;/U^kZpcL%LJ\5LJm-ALZ/SkIngRn73**MJFF          `
	Q_ = "90101191001109109000901000111000000100000010000000000"
	G_ = "----------------------------------------"
	F_ = "LET YOUR QUEST BEGIN"
	K_ = "REFLECTIONS STIR WITHIN"
	L_ = "REMEMBER ALADDIN IT WORKED FOR HIM"
	M_ = "DECIDE TO HAVE A LITTLE FUN AND "
	N_ = "THE TORCH BRIGHTENS"
	U_ = "YOU ANGER THE BIRD"
	W_ = "YOU CAN'T "
	P_ = "REMEMBER OLD TIMES"
	R_ = " WHICH FLIES YOU TO A REMOTE PLACE"
	S_ = "FALLING UNDER THE SPELL OF THE BOATMAN "
	T_ = "AND ARE TAKEN TO THE ISLAND OF SECRETS"
	J_ = "THE STONES ARE FIXED"
	O_ = "WHICH TAKES OMEGAN IN ITS CLAWS AND FLIES AWAY"
	V_ = "N??S??E??W??GO?GETTAKGIVDROLEAEATDRIRIDOPEPICCHOCHITAPBREFIGSTRATT"
	V_ = V_ + "HITKILSWISHEHELSCRCATRUBPOLREAEXAFILSAYWAIRESWAVINFXLOXSAQUI"
	Z_ = "APPEGGFLOJUGRAGPARTORPEBAXEROPSTACHICOAFLIHAMCANLOAMELBISMUS"
	Z_ = Z_ + "BOTWINSAPWATBOACHECOLSTOTRAVILLIQSWASAGBOOROOASAWRACLOOMESNA"
	Z_ = Z_ + "LOGSCAMEDNORSOUEASWESUP?DOWIN?OUT???"
	for I := 1; I <= W+1; I++ {
		L[I] = ASC(MID_(H_, I, 1)) - 32
		F[I] = ASC(MID_(Q_, I, 1)) - 48
	}
	H_ = "STONY WORDS"
}

// 4430
func FNR(Z int) int {
	return INT(RND_1()*float64(Z)) + 1
}

// 4440
func FNP(Z int) int {
	return POS()
}

// 4450
func FNS(Z int) int {
	return Y - INT((float64(E)/float64(C4) + .1))
}

var DATA_POINTER int
var DATA = []string{

	// 2870
	"4THE FURTHEST DEPTHS OF THE FOREST1001",
	"4THE DEPTHS OF THE MUTANT FOREST1000",
	"7A PATH OUT OF THE OVERGROWN DEPTHS1000",
	"6A CARNIVOROUS TREE1000",
	"4A CORRAL BENEATH THE CRIMSON CANYON1110",
	"7THE TOP OF A STEEP CLIFF1011",
	"4THE MARSH FACTORY1001",
	"4THE SLUDGE FERMENTATION VATS1110",
	"7THE UPPERMOST BATTLEMENTS1001",
	"4OMEGAN'S SANCTUM1110",
	"4SNELM'S LAIR0001",
	"2A DARK CAVE0000",
	"1BROKEN BRANCHES0100",
	"1A THICKET OF BITING BUSHES0000",
	"1A HUGE GLASSY STONE1110",
	"7THE EDGE OF CRIMSON CANYON0011",
	"4THE CLONE FACTORY0101",
	"4A CORRIDOR OF CLONE STORAGE CASKS1100",
	"7EDGE OF THE WELL0000",
	"4THE ROOM OF SECRET VISIONS1110",
	"4SNELM'S INNER CHAMBER0111",
	"3THE SOUTHERN EDGE OF THE FOREST0101",
	"7A LEAFY PATH1000",
	"3A FORK IN THE PATH0100",
	"7AN APPARENTLY UNCLIMBABLE ROCKY PATH1100",
	"7A LEDGE ATOP THE CRIMSON CANYON0010",
	"4A TALL ENTRANCE CHAMBER1101",
	"4A LOW PASSAGE WITH ARMS REACHING FROM THE WALLS1010",
	"7THE APPROACH TO THE WELL OF DESPAIR0001",
	"4A DIM CORRIDOR DEEP IN THE CASTLE1010",
	"4THE STAGNANT WATERS OF THE CRAWLING CREEK1001",
	"4A SHALLON POOL OFF THE CREEK1100",
	"7A LOG PIER, JUTTING OUT OVER THE CREEK0000",
	"4A STRETCH OF FEATURELESS DUNES1100",
	"1A GROUP OF TALL TREES1010",
	"7A NARROW LEDGE AT THE SUMMIT OF THE CANYON0011",
	"2A MONSTEROUS PORTAL IN THE CASTLE WALL0011",
	"4A CHAMBER INCHES DEEP WITH DUST0001",
	"4HERE1111",
	"2A CARVED ARCHWAY0010",
	"4A SMALL HUT IN THE LOG SETTLEMENT0111",
	"1A HUGE SPLIT-LOG TABLE1001",
	"4THE PORCH OF THE LOGMEN'S CABIN0110",
	"4GRANDPA'S SHACK1101",
	"3A CLEARING IN THE TREES BY A RICKETY SHACK0010",
	"4THE NEST OF A HUGE DACTYL0111",
	"6THE CASTLE OF DARK SECRETS BY TWO HUGE STONES0011",
	"4A ROOM LITTERED WITH BONES0111",
	"4THE CELL OF WHISPERED SECRETS0111",
	"4THE LIBRARY OF WRITTEN SECRETS0111",
	"4A REFUSE STREWN STOREROOM1111",
	"4THE LOGMEN'S HALL0000",
	"5A LOG BUILDING1000",
	"7A RUTTED HILLSIDE1100",
	"7A WINDSWEPT PLAIN AMONGST STONE MEGALITHS0100",
	"7THE STEPS OF AN ANCIENT PYRAMID1010",
	"7THE ISLAND OF SECRETS0111",
	"1A BROKEN MARBLE COLUMN1001",
	"7AN EXPANSE OF CRACKED, BAKED EARTH1100",
	"4A DESERTED ADOBE HUT1010",
	"4A LIVID GROWTH OF MAD ORCHIDS1011",
	"4A CORNER STREWN WITH BROKEN CHAIRS0111",
	"7THE BRIDGE NEAR TO A LOG SETTLEMENT0011",
	"1A CRUMBLING MASS OF PETRIFIED TREES1011",
	"3THE EDGE OF THE PYRAMID1101",
	"7THE ROOF OF THE ANCIENT PYRAMID0100",
	"3AN IMPASSABLE SPLIT IN THE PYRAMID1110",
	"7A BARREN BLASTED WASTELAND0001",
	"4AN EXPANSE OF BLEAK, BURNT LAND1100",
	"5A DELAPIDATED ADOBE HUT0110",
	"4THE HEART OF THE LILIES0101",
	"4THE MIDST OF THE LILIES1100",
	"3A RIVER'S EDGE BY A LOG BRIDGE0100",
	"3A PETRIFIED VILLAGE BY A RIVER CROWDED WITH LILIES0100",
	"4THE REMAINS OF A VILLAGE1100",
	"3THE ENTRANCE TO A PETRIFIED VILLAGE1100",
	"4A SWAMP MATTED WITH FIBROUS ROOTS1100",
	"2A VILLAGE OF HOLLOW STUMPS DEFYING THE SWAMP0100",
	"4A TUNNEL INTO ONE OF THE TREE STUMPS1100",
	"4A HOLLOW CHAMBER MANY METRES IN DIAMETER1110",

	// 3670
	"A SHINY APPLE",
	"A FOSSILISED EGG",
	"A LILY FLOWER",
	"AN EARTHENNARE JUG",
	"A DIRTY OLD RAG",
	"A RAGGED PARCHMENT",
	"A FLICKERING TORCH",
	"A GLISTENING PEBBLE",
	"A WOODMAN'S AXE",
	"A COIL OF ROPE",
	"A RUGGED STAFF",
	"A CHIP OF MARBLE",
	"A POLISHED COAL",
	"A PIECE OF FLINT",
	"A GEOLOGIST'S HAMMER",
	"A WILD CANYON BEAST",
	"A GRAIN LOAF",
	"A JUICY MELON",
	"SOME BISCUITS",
	"A GROWTH OF MUSHROOMS",
	"A BOTTLE OF WATER",
	"A FLAGON OF WINE",
	"A FLOWING SAP",
	"A SPARKLING FRESHWATER SPRING",
	"THE BOATMAN",
	"A STRAPPED OAK CHEST",
	"A FRACTURE IN THE COLUMN",
	"A MOUTH-LIKE OPENING",
	"AN OPEN TRAPDOOR",
	"A PARCHED, DESSICATED VILLAGER",
	"A STILL OF BUBBLING GREEN LIQUOR",
	"A TOUGH SKINNED SWAMPMAN",
	"THE SAGE OF THE LILIES",
	"WALL AFTER WALL OF EVIL BOOKS",
	"A NUMBER OF SOFTER ROOTS",
	"FIERCE LIVING STORM THAT FOLLOWS YOU",
	"MALEVOLENT WRAITHS WHO PUSH YOU TOWARD THE WELL",
	"HIS DREADED CLOAK OF ENTROPY",
	"OMEGAN THE EVIL ONE",
	"AN IMMENSE SNAKE WOUND AROUND THE HUT",
	"A GROUP OF AGGRESSIVE LOGMEN",
	"THE ANCIENT SCAVENGER",
	"MEDIAN",

	// 4100
	"BY", "FACING", "AT", "IN", "OUTSIDE", "BENEATH", "ON",
}

func READ() string {
	val := DATA[DATA_POINTER]
	DATA_POINTER++
	return val
}
