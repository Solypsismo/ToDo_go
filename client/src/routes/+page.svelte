<script>
    let tasks = [];

    const aggiungi = async (event) => {
        event.preventDefault();
        const form = event.target;

        const obj = Object.fromEntries(new FormData(form));

        const resJSON = await fetch("http://localhost:8080/api/task", {
            method: "POST"
        });

        const res = await resJSON.json();

        if(res.success) {
            tasks = res.tasks;
        }
    }
</script>

<div class="w-screen h-screen bg-slate-300 flex justify-center items-center gap-10">

    <div class="w-1/3 h-2/3 bg-accent p-5">
        <form method="POST" on:submit={aggiungi}>
            <input type="text" placeholder="Type here" name="task" class="input input-bordered w-full" />
            <div class="flex justify-center items-center p-3">
                <button class="btn btn-outline">Aggiungi</button>
            </div>
        </form>
    </div>

    <div class="w-1/3 h-2/3 bg-accent p-5">

    </div>


</div>