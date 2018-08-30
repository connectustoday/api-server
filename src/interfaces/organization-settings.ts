import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IOrganizationSettings extends IAccountSettings {

	@prop({ required: true })
	is_nonprofit: boolean;

}
