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

  if (data.skates == null) {
    data.skates = {};
  }

  $.each(data.skates, function(index, value) {
    skateData += '<tr>';
    skateData += '<td>' + value.type + '</td><td>' + value.startTime + '</td><td>' + value.endTime + '</td>';
    skateData += '</tr>';
  });

  $('table.co-skates tbody').html(skateData);
};

