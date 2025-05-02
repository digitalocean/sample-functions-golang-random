const actionEndpoint = '/functions/nodejs/say-it-sammy'

const getStatementButton = document.getElementById('get-sammy-statement-btn')
const statementText = document.getElementById('sammy-statement-text')

getStatementButton.addEventListener('click', (e) => {
    console.log('hello');
    const stmtPromise = getSammyStatement()
    stmtPromise.then(stmt => statementText.innerHTML = stmt["greeting"])
})

const getSammyStatement = () => {
    console.log(`fetching sammys thought from: ${actionEndpoint}`)
    return fetch(actionEndpoint)
        .then(response => response.json())
        .catch(e => console.error(e))
}
