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

type Result struct {
	// Rezultāta unikāls identifikators. Ja rezultāts ir par jau iepriekš iesūtītu izmeklējumu, tad šeit ir jāiesūta iepriekš saņemtais identifikators no response servisa par konkrēto rezultātu. Ja šis ir jauns rezultāts, lauks jāatstāj tukšs
	ResultId string `json:"resultId"`
	// Iesūtītā rezultāta versija (augoša vērtība pret iepriekšējo rezultātu) 
	VersionId int32 `json:"versionId"`

	ResultStatus int32 `json:"resultStatus"`
	// Datums un laiks, kurā ārstniecības persona apstiprināja testa rezultātu. Tiek sagaidīts laiks UTC laika zonā 
	ResultDate time.Time `json:"resultDate"`

	ResultType *ResultType `json:"resultType"`

	ResultContent string `json:"resultContent"`
	// Rezultāta mērvienība, ja kvantitatīvi dati
	ResultMetric string `json:"resultMetric,omitempty"`
	// Rezultāta references intervāls
	ResultReferenceValue string `json:"resultReferenceValue,omitempty"`
}