package email

var emailTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Good news! {{.Campground.Name}} is available</title>
	<style>
		table {
			width: 100%;
			border-collapse: collapse;
		}
		th, td {
			border: 1px solid #dddddd;
			text-align: left;
			padding: 8px;
		}
		tr:nth-child(even) {
			background-color: #f2f2f2;
		}
	</style>
</head>
<body>
	<h1>Good news! {{.Campground.Name}} is available</h1>
	<p>{{.Campground.Park}}, near {{.Campground.City}}</p>
	<p>Check-in: {{.CheckIn}}</p>
	<p>Check-out: {{.CheckOut}}</p>
	<table>
		<tr>
			<th>Campsite</th>
			<th>Loop</th>
			<th>Type</th>
			<th>Use</th>
			<th>Min People</th>
			<th>Max People</th>
			<th>Book Now</th>
		</tr>
		{{range .Campsites}}
		<tr>
			<td>{{.Name}}</td>
			<td>{{.Loop}}</td>
			<td>{{.Type}}</td>
			<td>{{.Use}}</td>
			<td>{{.MinPeople}}</td>
			<td>{{.MaxPeople}}</td>
			<td><a href="https://www.recreation.gov/camping/campsites/{{.ID}}">Book Now</a></td>
		</tr>
		{{end}}
	</table>
	<p style="font-size: small;">Brought to you by OpenCamp</p>
</body>
</html>
`
