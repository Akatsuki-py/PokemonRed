package mscript

import (
	"pokered/pkg/audio"
	"pokered/pkg/data/pkmnd"
	"pokered/pkg/data/txt"
	"pokered/pkg/data/worldmap"
	"pokered/pkg/event"
	"pokered/pkg/joypad"
	"pokered/pkg/menu"
	"pokered/pkg/pkmn"
	"pokered/pkg/sprite"
	"pokered/pkg/store"
	"pokered/pkg/text"
	"pokered/pkg/util"
	"pokered/pkg/widget"
)

func init() {
	txt.RegisterAsmText("OaksLabText1", func() string {
		if event.CheckEvent(event.EVENT_FOLLOWED_OAK_INTO_LAB_2) {
			if store.CurMapScript == 6 {
				return txt.OaksLabText40
			}
			return txt.OaksLabText41
		}
		return txt.OaksLabGaryText1
	})
	txt.RegisterAsmText("OaksLabLookAtCharmander", oaksLabLookAtCharmander)
	txt.RegisterAsmText("OaksLabLookAtSquirtle", oaksLabLookAtSquirtle)
	txt.RegisterAsmText("OaksLabLookAtBulbasaur", oaksLabLookAtBulbasaur)
	txt.RegisterAsmText("OaksYesNo", oaksYesNo)
	txt.RegisterAsmText("OaksLabMonChoiceMenu", oaksLabMonChoiceMenu)
	txt.RegisterAsmText("OaksLabText5", func() string {
		if event.CheckEvent(event.EVENT_FOLLOWED_OAK_INTO_LAB_2) {
			switch store.CurMapScript {
			case 6:
				return txt.OaksLabWhichMonText
			case 14:
				return txt.OaksLabRaiseMonText
			}
			return txt.OaksLabFightText
		}
		return txt.OaksLabGaryText1
	})
}

func oaksLabScript() {
	switch store.CurMapScript {
	case 0:
		oaksLabScript0()
	case 1:
		oaksLabScript1()
	case 2:
		oaksLabScript2()
	case 3:
		oaksLabScript3()
	case 4:
		oaksLabScript4()
	case 5:
		oaksLabScript5()
	case 6:
		oaksLabScript6()
	case 7:
		oaksLabScript7()
	case 8:
		oaksLabScript8()
	case 9:
		oaksLabScript9()
	case 14:
		oaksLabScript14()
	}
}

func oaksLabScript0() {
	if !event.CheckEvent(event.EVENT_OAK_APPEARED_IN_PALLET) {
		return
	}

	showObject(8)
	store.CurMapScript = 1
}

func oaksLabScript1() {
	oak := store.SpriteData[8]
	oak.DoubleSpd = false
	oak.Simulated = []uint{util.Up, util.Up, util.Up}
	store.CurMapScript = 2
}

func oaksLabScript2() {
	oak := store.SpriteData[8]
	if len(oak.Simulated) > 0 || oak.MovmentStatus == sprite.Movement {
		return
	}

	hideObject(8) // walking oak
	showObject(5) // staying oak
	store.CurMapScript = 3
	delayed3F = false
}

func oaksLabScript3() {
	if !delayed3F {
		store.DelayFrames = 3
		delayed3F = true
		return
	}

	p := store.SpriteData[0]
	p.Simulated = make([]uint, 8)
	for i := 0; i < 8; i++ {
		p.Simulated[i] = util.Up
	}

	store.CurMapScript = 4
}

func oaksLabScript4() {
	p := store.SpriteData[0]
	if len(p.Simulated) > 0 || p.MovmentStatus == sprite.Movement {
		return
	}

	event.UpdateEvent(event.EVENT_FOLLOWED_OAK_INTO_LAB, true)
	event.UpdateEvent(event.EVENT_FOLLOWED_OAK_INTO_LAB_2, true)

	blue := store.SpriteData[1]
	blue.Direction = util.Up

	audio.PlayDefaultMusic(worldmap.OAKS_LAB)

	store.CurMapScript = 5
}

var delay30FCtrInoaksLabScript5 = 0

func oaksLabScript5() {
	switch delay30FCtrInoaksLabScript5 {
	case 0:
		joypad.JoyIgnore = joypad.ByteToInput(0xfc)
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabText17, false)
		delay30FCtrInoaksLabScript5++
	case 1:
		store.DelayFrames = 20
		delay30FCtrInoaksLabScript5++
	case 2:
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabText18, false)
		delay30FCtrInoaksLabScript5++
	case 3:
		store.DelayFrames = 20
		delay30FCtrInoaksLabScript5++
	case 4:
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabText19, false)
		delay30FCtrInoaksLabScript5++
	case 5:
		store.DelayFrames = 20
		delay30FCtrInoaksLabScript5++
	case 6:
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabText20, false)
		delay30FCtrInoaksLabScript5++
		event.UpdateEvent(event.EVENT_OAK_ASKED_TO_CHOOSE_MON, true)
		joypad.JoyIgnore = joypad.ByteToInput(0x00)
		store.CurMapScript = 6
	case 7:
		store.DelayFrames = 20
		delay30FCtrInoaksLabScript5++
	}
}

func oaksLabScript6() {
	p := store.SpriteData[0]
	if p.MapYCoord != 6 {
		return
	}

	p.Direction = util.Up
	text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabText12, false)
	p.Simulated = []uint{util.Up}
	joypad.JoyIgnore = joypad.ByteToInput(0xfc)
	store.CurMapScript = 7
}

func oaksLabScript7() {
	p := store.SpriteData[0]
	if len(p.Simulated) > 0 || p.MovmentStatus == sprite.Movement {
		return
	}
	joypad.JoyIgnore = joypad.ByteToInput(0x00)
	store.CurMapScript = 6
}

func oaksLabLookAtCharmander() string {
	if store.Debug {
		widget.TargetMonID = pkmnd.CHARMANDER
		store.SetScriptID(store.WidgetStarterPokedexPage)
		store.Player.Starter = pkmnd.CHARMANDER
		store.Player.StarterName = pkmnd.Name(pkmnd.CHARMANDER)
		return ""
	}
	if event.CheckEvent(event.EVENT_GOT_STARTER) {
		return txt.OaksLabLastMonText
	}
	if !event.CheckEvent(event.EVENT_OAK_ASKED_TO_CHOOSE_MON) {
		return txt.OaksLabText39
	}

	widget.TargetMonID = pkmnd.CHARMANDER
	store.SetScriptID(store.WidgetStarterPokedexPage)
	store.Player.Starter = pkmnd.CHARMANDER
	store.Player.StarterName = pkmnd.Name(pkmnd.CHARMANDER)
	return ""
}

func oaksLabLookAtSquirtle() string {
	if event.CheckEvent(event.EVENT_GOT_STARTER) {
		return txt.OaksLabLastMonText
	}
	if !event.CheckEvent(event.EVENT_OAK_ASKED_TO_CHOOSE_MON) {
		return txt.OaksLabText39
	}

	widget.TargetMonID = pkmnd.SQUIRTLE
	store.SetScriptID(store.WidgetStarterPokedexPage)
	store.Player.Starter = pkmnd.SQUIRTLE
	store.Player.StarterName = pkmnd.Name(pkmnd.SQUIRTLE)
	return ""
}

func oaksLabLookAtBulbasaur() string {
	if event.CheckEvent(event.EVENT_GOT_STARTER) {
		return txt.OaksLabLastMonText
	}
	if !event.CheckEvent(event.EVENT_OAK_ASKED_TO_CHOOSE_MON) {
		return txt.OaksLabText39
	}

	widget.TargetMonID = pkmnd.BULBASAUR
	store.SetScriptID(store.WidgetStarterPokedexPage)
	store.Player.Starter = pkmnd.BULBASAUR
	store.Player.StarterName = pkmnd.Name(pkmnd.BULBASAUR)
	return ""
}

func oaksLabMonChoiceMenu() string {
	switch store.Player.Starter {
	case pkmnd.CHARMANDER:
		hideObject(2)
	case pkmnd.SQUIRTLE:
		hideObject(3)
	case pkmnd.BULBASAUR:
		hideObject(4)
	}

	store.CurMapScript = 8
	joypad.JoyIgnore = joypad.ByteToInput(0xfc)
	pkmn.AddPlayerPartyMon(store.Player.Starter, 5)

	return ""
}

func oaksLabScript8() {
	p := store.SpriteData[0]
	blue := store.SpriteData[1]
	switch store.Player.Starter {
	case pkmnd.CHARMANDER:
		store.Rival.Starter = pkmnd.SQUIRTLE
		store.Rival.StarterName = pkmnd.Name(pkmnd.SQUIRTLE)
		blue.Simulated = []uint{util.Down, util.Right, util.Right, util.Right}
		if p.MapYCoord == 4 {
			blue.Simulated = []uint{util.Down, util.Down, util.Right, util.Right, util.Right, util.Up}
		}
	case pkmnd.SQUIRTLE:
		store.Rival.Starter = pkmnd.BULBASAUR
		store.Rival.StarterName = pkmnd.Name(pkmnd.BULBASAUR)
		blue.Simulated = []uint{util.Down, util.Right, util.Right, util.Right, util.Right}
		if p.MapYCoord == 4 {
			blue.Simulated = []uint{util.Down, util.Down, util.Right, util.Right, util.Right, util.Right, util.Up}
		}
	case pkmnd.BULBASAUR:
		store.Rival.Starter = pkmnd.CHARMANDER
		store.Rival.StarterName = pkmnd.Name(pkmnd.CHARMANDER)
		blue.Simulated = []uint{util.Down, util.Right, util.Right}
	}

	store.CurMapScript = 9
}

var script9Phase int

func oaksLabScript9() {
	switch script9Phase {
	case 0:
		blue := store.SpriteData[1]
		if len(blue.Simulated) > 0 || blue.MovmentStatus == sprite.Movement {
			return
		}

		blue.Direction = util.Up
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabRivalPickingMonText, false)
		script9Phase++
	case 1:
		switch store.Rival.Starter {
		case pkmnd.CHARMANDER:
			hideObject(2)
		case pkmnd.SQUIRTLE:
			hideObject(3)
		case pkmnd.BULBASAUR:
			hideObject(4)
		}
		pkmn.AddRivalPartyMon(store.Rival.Starter, 5)
		text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabRivalReceivedMonText, false)
		script9Phase++
	case 2:
		event.UpdateEvent(event.EVENT_GOT_STARTER, true)
		joypad.JoyIgnore = joypad.ByteToInput(0x00)
		store.CurMapScript = 14
		script9Phase++
	}
}

func oaksYesNo() string {
	store.SetScriptID(store.TwoOptionMenu)
	store.PushOtScript(func() {
		switch store.TwoOptionResult {
		case 0:
			text.DoPrintTextScript(text.TextBoxImage, txt.OaksLabReceivedMonText, true)
		case 1:
			text.CloseTextBox()
			store.Player.Starter = 0
			store.CurMapScript = 6
			return
		}
	})
	menu.NewYesNoMenu()
	return ""
}

func oaksLabScript14() {
}
