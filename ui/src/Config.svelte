<script>
    import axios from 'axios';

    let addressEditMode = false
    let newAddress = ''
    async function editAdress(address) {
        let result = await axios.put('/config/higheraddress', address).then(response => response.data)
        alert(result)
        addressEditMode = false
    }
</script>