export async function load() {
    const resJSON = await fetch("http://backend:8080/api/tasks", {
        method: "GET"
    });

    if(resJSON.ok) {
        let res = await resJSON.json();
        if(res == null) res = []
        return {"tasks" : res};
    }else {
        return {"success" : false};
    }
}