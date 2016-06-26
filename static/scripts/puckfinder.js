function updateTable(id, tableId) {
  $.ajax({
    url: "/api/rinks/" + id,
    dataType: "json",
    success: function(data) {
      updateHTMLForTable(tableId, data);
    },
    error: function(msg) {
      alert(msg.statusText);
    }
  });
}

function updateHTMLForTable(tableId, data) {
  var skateData = '';
  $.each(data.skates, function(index, value) {
    skateData += '<tr>';
    skateData += '<td>' + value.type + '</td><td>' + value.startTime + '</td><td>' + value.endTime + '</td>';
    skateData += '</tr>';
    console.log("skateData is" + skateData);
  });

  $('table.co-skates tbody').html(skateData);
};
