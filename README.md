A simple rhythm generator/manipulator toolbox.


# Development

1.  Download the Dockerfile and create an image (takes some time):

    `docker build -t rhythmtool-env .`

2.  Run it in interactive terminal mode (the port is only for viewing the website):

    `docker run -it --name rhythmtool-env-1 -p 8081:8081 rhythmtool-env`

3.  Check if the GUI works, by opening a chrome browser on the host system and visiting
    `localhost:8081/app/`. (admin/admin)

4.  Run the test and generate the docs in the docker container and exit it:

    ```sh
    cd guitest
    go generate

    exit
    ```

5.  If the test passed, you can copy the doc to your host:

    `docker cp rhythmtool-env-1:/go/src/github.com/peter-mueller/rhythmtool/guitest/doc.html doc.html`
