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
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/lab-api/",
		Index,
	},

	Route{
		"GetDZSforPerson",
		strings.ToUpper("Get"),
		"/lab-api/api/DZSdata/{PersonId}",
		GetDZSforPerson,
	},

	Route{
		"GetDZSforResult",
		strings.ToUpper("Get"),
		"/lab-api/api/DZSdata/{ResultId}",
		GetDZSforResult,
	},

	Route{
		"GetLabRecordsStatus",
		strings.ToUpper("Get"),
		"/lab-api/api/LabRecordsStatus/{PostId}",
		GetLabRecordsStatus,
	},

	Route{
		"LabRecordsDelete",
		strings.ToUpper("Delete"),
		"/lab-api/api/LabRecords/{RecordId}",
		LabRecordsDelete,
	},

	Route{
		"LabRecordsGet",
		strings.ToUpper("Get"),
		"/lab-api/api/LabRecords/{RecordId}",
		LabRecordsGet,
	},

	Route{
		"LabRecordsPost",
		strings.ToUpper("Post"),
		"/lab-api/api/LabRecords",
		LabRecordsPost,
	},

	Route{
		"LabRecordsPut",
		strings.ToUpper("Put"),
		"/lab-api/api/LabRecords/{RecordId}",
		LabRecordsPut,
	},

	Route{
		"GetLabRecordsBatchStatus",
		strings.ToUpper("Get"),
		"/lab-api/api/LabRecordsBatch/{Batchid}",
		GetLabRecordsBatchStatus,
	},

	Route{
		"LabRecordsBatchDeleteLabRecordsBatch",
		strings.ToUpper("Delete"),
		"/lab-api/api/LabRecordsBatch/{Batchid}",
		LabRecordsBatchDeleteLabRecordsBatch,
	},

	Route{
		"LabRecordsBatchSaveLabRecordsBatch",
		strings.ToUpper("Post"),
		"/lab-api/api/LabRecordsBatch",
		LabRecordsBatchSaveLabRecordsBatch,
	},

	Route{
		"LabRecordsBatchUpdateLabRecordsBatch",
		strings.ToUpper("Put"),
		"/lab-api/api/LabRecordsBatch/{Batchid}",
		LabRecordsBatchUpdateLabRecordsBatch,
	},
}
