import React from "react";

class AdminProductPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            message: "",
            id: window.location.pathname.split("/").pop(),
            thumbnail: "",
            // TODO: Add image previews
            images: []
        };
        this.submit = this.submit.bind(this)
        this.thumbnailUpload = this.thumbnailUpload.bind(this)
    }

    componentDidMount() {
        return
    }

    submit(event) {
        event.preventDefault()
        const data = new FormData(event.target)
        const options = {
            method: "PUT",
            body: data
        };
        fetch("/api/product/" + this.state.id, options)
            .then(res => res.json())
            .then(
                (result) => {
                    if (result.message) {
                        this.setState({
                            message: result.message
                        });
                    }
                },
                (error) => {
                    this.setState({
                        message: error 
                    });
                }
            )
    }

    thumbnailUpload(event) {
        this.setState({
            thumbnail: URL.createObjectURL(event.target.files[0])
        })
    }

    render() {
        return (
            <div className="container-fluid">
                <form className="col-sm-10 m-auto p-2" onSubmit={this.submit}>
                    <div className="form-group">
                        <label>Name</label>
                        <input className="form-control" type="text" name="name" />
                    </div>
                    <div className="form-group">
                        <label>Identifier</label>
                        <input className="form-control" type="text" name="identifier" />
                        <small className="form-text text-muted">Unique identifer: No spaces, No special characters</small>
                    </div>
                    <div className="form-check">
                        <input className="form-check-input" type="checkbox" name="public" />
                        <label>Public</label>
                    </div>
                    <div className="form-group">
                        <label>Thumbnail</label>
                        <input className="form-control-file" type="file" name="thumbnail" onChange={this.thumbnailUpload} />
                        <img src={this.state.thumbnail} className="mini-image my-2" />
                    </div>
                    <div className="form-group">
                        <label>Images</label>
                        <input className="form-control-file" type="file" name="images" multiple />
                    </div>
                    <div className="form-group">
                        <label>Description</label>
                        <textarea className="form-control" rows="3" name="description" />
                    </div>
                    <div className="form-group">
                        <label>Price</label>
                        <input className="form-control" type="number" name="price" />
                    </div>
                    <div className="form-group">
                        <label>Discount</label>
                        <input className="form-control" type="number" step="0.01" min="0" max="1" name="discount" />
                        <small className="form-text text-muted">Discount multiplier applied to price</small>
                    </div>
                    <button type="submit" className="btn btn-primary float-right">Submit</button>
                </form>
            </div>
        );
    }
}

export default AdminProductPage;
