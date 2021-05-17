/*
 * Lab COVID-19 result API
 *
 * Dokumentācija laboratoriju COVID-19 testa datu API iesūtīšanai. Datu iesūtīšanas mērķis ir saņemt nepieciešamo informāciju par COVID-19 testiem, ko tālāk izmantotu vairākiem mērķiem: DZS (Digitālā Zaļā Sertifikāta) izveidei, VMNVD (LR Veselības ministrijas Nacionālā veselības dienesta) vajadzībām, SPKC (Slimību profilakses un kontroles centra) vajadzībām un VI (Veselības Inspekcijas) vajadzībām. Izveidotās datu struktūras ir specifiskas COVID-19 vajadzībām, kā arī to izveidē ir mēģināts pēc iespējas saglabāt un atkārtoti izmantot EVK klasifikatorus un FHIR datu struktūru nosaukumus atbilstoši: https://www.hl7.org/fhir/diagnosticreport.html    COVID-19 Diagnostikas pārskata pievienošana notiek asinhronā režīmā (primārā atgrieztā atbilde vienmēr ir HTTP 202 (Accepted)). Atbildē tiek atgriezts konkrētā izsaukuma statusa pārbaudes URL, kuru izmantojot organizācija var pieprasīt iesūtītā pierasījuma verifikācijas statusu. Verifikācijas statusa pārbaude jāveic, izmantojot Ekponenciālo atkāpšanās pieeju: https://en.wikipedia.org/wiki/Exponential_backoff 
 *
 * API version: 1.0
 * Contact: zzdats@zzdats.lv
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Examination struct {
	// Vizītes unikālais identifikators Laboratorijas sistēmā
	EncounterId string `json:"encounterId"`
	// Izmeklējuma veikšanas datums un laiks. Tiek sagaidīts laiks UTC zonā
	TestDate time.Time `json:"testDate"`
	// Izmeklējuma mērķis. Pašreiz konstante COVID-19 
	TestSubject string `json:"testSubject"`
	// Izmeklējuma veids NAAT/RAT (provizoriski no https://loinc.org/sars-cov-2-and-covid-19/ )
	TestType string `json:"testType"`

	TestName string `json:"testName,omitempty"`

	TestManufacturer string `json:"testManufacturer,omitempty"`
	// Materiāla unikālais identifikators Laboratoriju sistēmā
	SampleId string `json:"sampleId"`

	SampleType int32 `json:"sampleType"`
	// Ārstniecības persona, kura atbild par testa rezultātu, Veselības inspecijas piešķirtais identifikators. Iesūta identifikatoru no 2.37 un pēc id18 koda pārbauda pret 2.1 id4 
	ResponsibleForTestResult string `json:"responsibleForTestResult"`
	// Ārstniecības personas, kura atbild par testa rezultātu telefona numurs, ietverot valsts kodu un 00 prefiksu
	ResponsiblePersonPhone string `json:"responsiblePersonPhone"`

	OrdererOfTest *OrdererOfTest `json:"ordererOfTest"`
	// Testējamās grupas nosaukums 
	TestGroup string `json:"testGroup,omitempty"`
	// Testēšanas organizatora kontaktinformācija brīvā formā (vārds, uzvārds, telefons). SPKC izmantošanai 
	TestOrganizer string `json:"testOrganizer,omitempty"`

	ReferralType int32 `json:"referralType,omitempty"`
	// Nosūtījuma numurs (aizpildāms par papīra nosūtījumu un elektronisko nosūtījumu)
	Referral string `json:"referral,omitempty"`

	ReferralMedicalPerson string `json:"referralMedicalPerson,omitempty"`

	PayerForTest int32 `json:"payerForTest,omitempty"`
	// Indentificēts COVID-19 infekcijas celms. Provizoriski atbilstošs LOINC pre-release klasifikatoram: 96741-4
	Cause string `json:"cause,omitempty"`
	// Datu ievadītāja personas kods
	Operator string `json:"operator"`

	ResultList []Result `json:"resultList"`
}
