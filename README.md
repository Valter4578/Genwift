![swift logo](https://eclipsesource.com/wp-content/uploads/2014/06/Apple_Swift_Logo-150x150.png)
# Genwift
Simple utility written on Go to create templates for modules in swift 
# Usage 
run ```go build``` and copy binary file to your root directory(for example into your user's directory).

For create MVP module with assembly, presenter, view controller, protocols for view and presenter you need to write module name as argument.
```~/genwift ModuleName MVP```

# Files structure 
```~/genwift Login MVP```

#### LoginViewController.swift 
```swift 
class LoginViewController: UIViewController { 
	// MARK:- Dependencies 
	var presenter: LoginOutput! 

	// MARK:- Lifecycle 
	override func viewDidLoad() {
		super.viewDidLoad()

	}

	// MARK:- <#T## mark's name#>
}

// MARK: LoginInput Implementation
extension LoginViewController: LoginInput {

}
```

#### LoginPresenter.swift 

```swift 
class LoginPresenter: LoginOutput {
	// MARK:- Dependencies 
	weak var view: LoginInput!
	// MARK:- <#T## mark's name #> 
	<#T##Presenter's code#>

	// MARK:- Initializers 
	init(view: LoginInput) {
		self.view = view
	}
}
```

#### LoginInput.swift 
```swift 
protocol LoginInput: class {
	<#T##Input's protocol code#>
}

```

#### LoginOutput.swift 
```swift
protocol LoginOutput {
	<#T##Output's protocol code#>
}
```

#### LoginAssembly.swift 

```swift 
class LoginAssembly {
	static func configureModule() -> LoginViewController {
		let view = LoginViewController()
		let presenter = LoginPresenter(view: view)
		view.presenter = presenter 
		return view 
	}
}
```