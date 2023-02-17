package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/
func GetFunFacts(about string) []string

type FunFacts struct {
  Sun [2]string {
    "Temperatur i Solens kjerne er ",
    "Temperatur på ytre lag av Solen er ",
  }
  Luna [2]string{
    "Temperatur på Månens overflate om natten er ",
    "Temperatur på Månens overflate om dagen er ",
  }
  Terra [3]string{
    "Høyeste temperatur målt på Jordens overflate er ",
    "Laveste temperatur målt på Jordens overflate er ",
    "Temperatur i Jordens indre kjerne 	 er ",
  }
}