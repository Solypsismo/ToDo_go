<script>
    export let data;

    const aggiungi = async (event) => {
        event.preventDefault();
        const form = event.target;

        const obj = Object.fromEntries(new FormData(form));
        const arr = [obj]
        const resJSON = await fetch("http://localhost:8080/api/task", {
            method: "POST",
            body: JSON.stringify(arr)
        });

        if(resJSON.ok) {
            let lista = [...tasks];
            lista.push(obj);
            tasks = lista
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
                <div class=" bg-green-600 w-full rounded-lg p-5 text-purple-700 font-bold mb-2">
                    <div>{task.title}</div>
                </div>
            {/each}
        {/if}
    </div>


</div>