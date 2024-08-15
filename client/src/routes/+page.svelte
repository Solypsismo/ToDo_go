<script>
    export let data;

    function updateStatusToDone(lis, id) {
        for (let i = 0; i < lis.length; i++) {
            if (lis[i].id === id) {
            lis[i].status = "done";
            return lis
            }
        }
        return null;
    }

    const aggiungi = async (event) => {
        event.preventDefault();
        const form = event.target;

        const obj = Object.fromEntries(new FormData(form));
        const resJSON = await fetch("http://localhost:8080/api/tasks", {
            method: "POST",
            body: JSON.stringify(obj)
        });

        if(resJSON.ok) {
            let res = await resJSON.json()
            let lista = [...tasks];
            obj.id = res.id
            lista.push(obj);
            tasks = lista
            console.log(tasks)
        }
    }

    const rimuovi = async ( event ) => {
        const div = event.target.parentElement.parentElement
        const id = div.dataset.taskId

        const resJSON = await fetch("http://localhost:8080/api/tasks", {
            method: "DELETE",
            body: JSON.stringify({id})
        })

        if(resJSON.ok) {
            div.remove()
        }else {
            alert("Undefined error")
        }
    }

    const aggiorna = async ( event ) => {
        const div = event.target.parentElement.parentElement
        const id = div.dataset.taskId

        const resJSON = await fetch("http://localhost:8080/api/tasks", {
            method: "PUT",
            body: JSON.stringify({id})
        })

        if(resJSON.ok) {
            let ris = updateStatusToDone(tasks, id);
            tasks = ris;            
        }else {
            alert("Undefined error")
        }
    }

    $: tasks = data.tasks;

</script>

<div class="w-screen h-screen bg-slate-300 flex justify-center items-center gap-10">

    <div class="w-1/3 h-2/3 bg-slate-800 p-5">
        <form method="POST" on:submit={aggiungi}>
            <input type="text" placeholder="Type here" name="title" class="input input-bordered w-full" />
            <div class="flex justify-center items-center p-3">
                <button class="btn">Aggiungi</button>
            </div>
        </form>
    </div>

    <div class="w-1/3 h-2/3 bg-slate-800 p-5 overflow-scroll">
        {#if tasks.length == 0}
            <div>
                Non sono presenti task
            </div>
        {:else}
            {#each tasks as task}
                {#if task.status == "done"}
                    <div class="bg-gray-500 w-full rounded-lg p-5 text-purple-700 font-bold text mb-2 line-through">
                        {task.title}
                    </div>
                {:else}
                    <div class=" bg-green-600 w-full rounded-lg p-5 text-purple-700 font-bold mb-2 flex justify-between items-center" data-task-id={task.id}>
                        <div>{task.title}</div>
                        <div class="flex gap-1">
                            <button class="btn btn-warning" on:click={rimuovi}>X</button>
                            <button class="btn btn-success" on:click={aggiorna}>DONE</button>
                        </div>
                    </div>
                {/if}
            {/each}
        {/if}
    </div>


</div>