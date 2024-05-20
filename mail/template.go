package mail

const (
	OrderConfirmationTemplate = `
		<p>Merhaba {{ (index . 0).UserName }}</p>
		<p>Siparişiniz başarıyla alınmıştır,</p>
		<p>Sipariş No: {{ (index . 0).OrderId}}</p>
		<p>Sipariş Tutarı: {{ (index . 0).OrderPrice}}</p>
		<p>Sipariş Detayı:</p>
		<figure class="table">
			<table>
				<tbody>
					<tr>
						<td>No.</td>
						<td>Kitap Adı</td>
						<td>Fiyat</td>
						<td>Adet</td>
						<td>Tutar</td>
					</tr>
					{{range .}} 
					<tr>
						<td>&nbsp;</td>
						<td>{{.ItemName}}</td>
						<td>{{.ItemPrice}}</td>
						<td>{{.ItemQty}}</td>
						<td>{{.ItemTotalPrice}}</td>
					</tr>
					{{end}} 
				</tbody>
			</table>
		</figure>
`
)
