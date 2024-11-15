document.addEventListener('DOMContentLoaded', function () {
    const parameterForm = document.getElementById('parameter-form');
    const parametersList = document.getElementById('parameters-list');
    const relationshipForm = document.getElementById('relationship-form');
    const relationshipsList = document.getElementById('relationships-list');
    const partnerForm = document.getElementById('partner-form');
    const partnersList = document.getElementById('partners-list');

    parameterForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const name = document.getElementById('param-name').value;
        const type = document.getElementById('param-type').value;
        const defaultValue = document.getElementById('param-default').value;

        const parameter = {
            Name: name,
            Type: type,
            DefaultValue: defaultValue
        };

        fetch('/api/parameters', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(parameter)
        })
        .then(response => response.json())
        .then(data => {
            loadParameters();
            parameterForm.reset();
        })
        .catch(error => console.error('Error:', error));
    });

    function loadParameters() {
        fetch('/api/parameters', {
            method: 'GET'
        })
        .then(response => response.json())
        .then(data => {
            parametersList.innerHTML = '';
            data.forEach(parameter => {
                const li = document.createElement('li');
                li.textContent = `Name: ${parameter.Name}, Type: ${parameter.Type}, Default Value: ${parameter.DefaultValue}`;
                parametersList.appendChild(li);
            });
        })
        .catch(error => console.error('Error:', error));
    }

    relationshipForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const type = document.getElementById('rel-type').value;
        const paramIDs = document.getElementById('rel-params').value.split(',').map(Number);

        const relationship = {
            Type: type,
            Parameters: paramIDs.map(id => ({ ID: id }))
        };

        fetch('/api/relationships', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(relationship)
        })
        .then(response => response.json())
        .then(data => {
            loadRelationships();
            relationshipForm.reset();
        })
        .catch(error => console.error('Error:', error));
    });

    function loadRelationships() {
        fetch('/api/relationships', {
            method: 'GET'
        })
        .then(response => response.json())
        .then(data => {
            relationshipsList.innerHTML = '';
            data.forEach(relationship => {
                const li = document.createElement('li');
                li.textContent = `Type: ${relationship.Type}, Parameters: ${relationship.Parameters.map(p => p.ID).join(', ')}`;
                relationshipsList.appendChild(li);
            });
        })
        .catch(error => console.error('Error:', error));
    }

    partnerForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const name = document.getElementById('partner-name').value;
        const paramIDs = document.getElementById('partner-params').value.split(',').map(Number);

        const partner = {
            Name: name,
            Parameters: paramIDs.map(id => ({ ID: id }))
        };

        fetch('/api/partners', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(partner)
        })
        .then(response => response.json())
        .then(data => {
            loadPartners();
            partnerForm.reset();
        })
        .catch(error => console.error('Error:', error));
    });

    function loadPartners() {
        fetch('/api/partners', {
            method: 'GET'
        })
        .then(response => response.json())
        .then(data => {
            partnersList.innerHTML = '';
            data.forEach(partner => {
                const li = document.createElement('li');
                li.textContent = `Name: ${partner.Name}, Parameters: ${partner.Parameters.map(p => p.ID).join(', ')}`;
                partnersList.appendChild(li);
            });
        })
        .catch(error => console.error('Error:', error));
    }

    loadParameters();
    loadRelationships();
    loadPartners();
});
