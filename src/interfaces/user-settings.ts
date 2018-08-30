import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IUserSettings extends IAccountSettings {

	@prop({ required: true })
	is_full_name_visible: boolean;

	@prop({ required: true })
	blocked_users: Array<string>;

}
