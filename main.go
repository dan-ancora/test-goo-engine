package main

import (
	"bytes"
	"fmt"
	"github.com/ViajerosAdoquier/teamamerica"
	//"github.com/dan-ancora/prototeamamerica"
	"google.golang.org/appengine" // Required external App Engine library
	"net/http"
	"strings"
)

var (
	clientTeamAmerica = teamamerica.New("XMLSMAY", "M3WgnuOV", "https://javatest.teamamericany.com:8443/TADoclit/services/TADoclit")
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, "<h1>Hello, Google Api Engine!</h1>")
}

func getCitiesList(w http.ResponseWriter, r *http.Request) {

	response, mesaj, err := clientTeamAmerica.ListCities(r)
	if err != nil || response == nil {
		w.Write([]byte(fmt.Sprintf("Error: %s<hr>%s", err, mesaj)))
	} else {

		//return result: for testing just print cities as html table
		var xhtml bytes.Buffer

		xhtml.WriteString(`
		<html>
			<body>
					<div style="height: 70%; overflow-y: auto">		   
						<table>
							<tr>
								<th>City Code</th>
								<th>City Name<br><i><small>show Hotels</small></i></th>
								<th>Country</th>
								<th>Display Group</th>
							</tr>
			`)

		for _, v := range response.Cities {
			xhtml.WriteString(fmt.Sprintf(`
							<tr>
								<td>%s</td>
								<td><a href="/vendor_list/%s">%s</a></td>
								<td>%s</td>
								<td>%s</td>
							</tr>
				`,
				v.GetCode(),
				v.GetCode(), v.GetName(),
				v.GetCountry(),
				v.GetDisplaygroup()))
		}

		xhtml.WriteString(`
						</table>
					</div>

					<h3><a href="/">Back to Home</a></h3>
				</body>
			</html>
			`)

		xhtml.WriteString(mesaj)
		w.Write([]byte(xhtml.String()))
	}
}

func getVendorListCity(w http.ResponseWriter, r *http.Request) {

	cale := strings.Split(r.RequestURI, "/")
	if len(cale) < 3 {
		w.Write([]byte("Missing parameters"))
		return
	}

	//w.Write([]byte(cale[2] + "/" + cale[3]))

	response, mesaj, err := clientTeamAmerica.ListVendor(r, cale[2], "Hotel")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s\n%s", err, mesaj)))
	} else {
		//return result: for testing just print vendors as html table
		var xhtml bytes.Buffer

		xhtml.WriteString(fmt.Sprintf(`
		<html>
		<body>
			<h2>HOTELS</h2>
			<table>
			<tr>
				<th>Vendor ID</th>
				<th>Vendor Name</th>
				<th>City</th>
			</tr>
				`))

		for _, v := range response.Vendors {
			xhtml.WriteString(fmt.Sprintf(`
			<tr>
				<td>%v
				<td><a href="/product_info_v2/%v">%s</a>
				<td>%v
			</tr>`,
				v.GetId(), v.GetId(), v.GetName(), v.GetCity()))
		}

		xhtml.WriteString(`
			</table>
		</body>
		</html>
				`)
		w.Write([]byte(xhtml.String()))
	}

	response, mesaj, err = clientTeamAmerica.ListVendor(r, cale[2], "Service")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s\n%s", err, mesaj)))
	} else {
		//return result: for testing just print vendors as html table
		var xhtml bytes.Buffer

		xhtml.WriteString(fmt.Sprintf(`
		<html>
		<body>
			<h2>SERVICES</h2>
			<table>
			<tr>
				<th>Vendor ID</th>
				<th>Vendor Name</th>
				<th>City</th>
			</tr>
				`))

		for _, v := range response.Vendors {
			xhtml.WriteString(fmt.Sprintf(`
			<tr>
				<td>%v
				<td>%s
				<td>%v
			</tr>`,
				v.GetId(), v.GetName(), v.GetCity()))
		}

		xhtml.WriteString(`
			</table>
		</body>
		</html>
				`)
		w.Write([]byte(xhtml.String()))
	}

	xhtml, err2 := clientTeamAmerica.ListPickUpLocations(r, cale[2])
	if err2 != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s\n%s", err2, xhtml)))
	} else {
		w.Write([]byte(xhtml))
	}
}

func getRoomType(w http.ResponseWriter, r *http.Request) {
	response, err := clientTeamAmerica.ListRoomType(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s\n%s", err, response)))
	} else {
		w.Write([]byte(response))
	}
}

func getMealPlan(w http.ResponseWriter, r *http.Request) {
	response, err := clientTeamAmerica.ListMealPlan(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s\n%s", err, response)))
	} else {
		w.Write([]byte(response))
	}
}

func getProductInfoV2(w http.ResponseWriter, r *http.Request) {

	cale := strings.Split(r.RequestURI, "/")
	if len(cale) < 2 {
		w.Write([]byte("Missing parameters"))
		return
	}

	response, err := clientTeamAmerica.ProductInfov2(r, "", cale[2])
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s\n%s", err, response)))
	} else {
		w.Write([]byte(response))
	}
}

func getPriceSearch(w http.ResponseWriter, r *http.Request) {

	params := []byte(`
	{  
		"checkin":{
	   
		   "day":22,
		   "month":11,
		   "year":2018
		},
		"checkout":{
	   
		   "day":23,
		   "month":11,
		   "year":2018
		},
		"rooms":[  
		   {  
			  "adults":2
		   }
		],
		"lang":"es",
		"hotels":[  
		   {  
			  "id":3,
			  "provider_code":"15975"
		   },
		   {  
			  "id":3,
			  "provider_code":"14700"
		   }
		],
		"site":"totembo"
	 }
	`)

	cale := strings.Split(r.RequestURI, "/")
	if len(cale) < 2 {
		w.Write([]byte("Missing parameters"))
		return
	}

	response, err := clientTeamAmerica.PriceSearch(r, params)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s\n%s", err, response)))
	} else {
		w.Write([]byte(response))
	}
}

func main() {
	http.HandleFunc("/city_list", getCitiesList)

	http.HandleFunc("/vendor_list/", getVendorListCity)

	http.HandleFunc("/room_type", getRoomType)

	http.HandleFunc("/meal_plan", getMealPlan)

	http.HandleFunc("/product_info_v2/", getProductInfoV2)

	http.HandleFunc("/search_price/test", getPriceSearch)

	http.HandleFunc("/", indexHandler)

	appengine.Main() // Starts the server to receive requests
}
