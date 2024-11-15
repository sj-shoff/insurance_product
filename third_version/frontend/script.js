const API_URL = 'http://localhost:8000/api';

// Helper function to fetch data
async function fetchData(endpoint) {
    const response = await fetch(`${API_URL}/${endpoint}`);
    return await response.json();
}

// Helper function to post data
async function postData(endpoint, data) {
    const response = await fetch(`${API_URL}/${endpoint}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
    return await response.json();
}

// Helper function to delete data
async function deleteData(endpoint, id) {
    const response = await fetch(`${API_URL}/${endpoint}/${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    return await response.json();
}

// Load parameters
async function loadParameters() {
    const parameters = await fetchData('parameters');
    const parametersList = document.getElementById('parametersList');
    parametersList.innerHTML = '';
    parameters.forEach(parameter => {
        const li = document.createElement('li');
        li.textContent = `${parameter.Name} (${parameter.Type})`;
        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Delete';
        deleteButton.onclick = () => deleteParameter(parameter.ID);
        li.appendChild(deleteButton);
        parametersList.appendChild(li);
    });
}

// Create parameter
document.getElementById('createParameterForm').onsubmit = async function(event) {
    event.preventDefault();
    const formData = new FormData(this);
    const parameter = {
        Name: formData.get('name'),
        Type: formData.get('type'),
        DefaultValue: formData.get('defaultValue')
    };
    await postData('parameters', parameter);
    this.reset();
    loadParameters();
};

// Delete parameter
async function deleteParameter(id) {
    await deleteData('parameters', id);
    loadParameters();
}

// Load relationships
async function loadRelationships() {
    const relationships = await fetchData('relationships');
    const relationshipsList = document.getElementById('relationshipsList');
    relationshipsList.innerHTML = '';
    relationships.forEach(relationship => {
        const li = document.createElement('li');
        li.textContent = `${relationship.Type}`;
        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Delete';
        deleteButton.onclick = () => deleteRelationship(relationship.ID);
        li.appendChild(deleteButton);
        relationshipsList.appendChild(li);
    });
}

// Create relationship
document.getElementById('createRelationshipForm').onsubmit = async function(event) {
    event.preventDefault();
    const formData = new FormData(this);
    const relationship = {
        Type: formData.get('type')
    };
    await postData('relationships', relationship);
    this.reset();
    loadRelationships();
};

// Delete relationship
async function deleteRelationship(id) {
    await deleteData('relationships', id);
    loadRelationships();
}

// Load partners
async function loadPartners() {
    const partners = await fetchData('partners');
    const partnersList = document.getElementById('partnersList');
    partnersList.innerHTML = '';
    partners.forEach(partner => {
        const li = document.createElement('li');
        li.textContent = `${partner.Name}`;
        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Delete';
        deleteButton.onclick = () => deletePartner(partner.ID);
        li.appendChild(deleteButton);
        partnersList.appendChild(li);
    });
}

// Create partner
document.getElementById('createPartnerForm').onsubmit = async function(event) {
    event.preventDefault();
    const formData = new FormData(this);
    const partner = {
        Name: formData.get('name')
    };
    await postData('partners', partner);
    this.reset();
    loadPartners();
};

// Delete partner
async function deletePartner(id) {
    await deleteData('partners', id);
    loadPartners();
}

// Load initial data
loadParameters();
loadRelationships();
loadPartners();